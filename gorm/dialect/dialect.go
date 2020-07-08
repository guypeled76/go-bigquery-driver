package dialect

import (
	"fmt"
	_ "github.com/guypeled76/go-bigquery-driver/driver"
	"github.com/guypeled76/go-bigquery-driver/processor"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const Name = "bigquery"

type bigQueryDialect struct {
	db gorm.SQLCommon
}

func (b *bigQueryDialect) GetName() string {
	return Name
}

func (b *bigQueryDialect) SetDB(db gorm.SQLCommon) {
	b.db = db
}

func (b *bigQueryDialect) BindVar(i int) string {
	return "?"
}

func (b *bigQueryDialect) Quote(key string) string {
	return fmt.Sprintf("`%s`", key)
}

func (b *bigQueryDialect) DataTypeOf(field *gorm.StructField) string {

	var dataValue, sqlType, _, additionalType = parseFieldStructForDialect(field, b)

	if sqlType == "" {
		switch dataValue.Kind() {
		case reflect.Bool:
			sqlType = "BOOL"
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
			if b.fieldCanAutoIncrement(field) {
				sqlType = "INT64 AUTO_INCREMENT"
			} else {
				sqlType = "INT64"
			}
		case reflect.Int64, reflect.Uint64:
			if b.fieldCanAutoIncrement(field) {
				sqlType = "INT64 AUTO_INCREMENT"
			} else {
				sqlType = "INT64"
			}
		case reflect.Float32, reflect.Float64:
			sqlType = "FLOAT64"
		case reflect.String:
			sqlType = "STRING"

		case reflect.Struct:
			if _, ok := dataValue.Interface().(time.Time); ok {
				sqlType = "TIMESTAMP"
			} else {
				sqlType = b.dataTypeOfStruct(field.Struct)
			}
		case reflect.Array:
			if _, ok := dataValue.Interface().(uuid.UUID); ok {
				sqlType = "STRING"
			} else {
				sqlType = fmt.Sprintf("ARRAY<%s>", b.dataTypeOfStruct(field.Struct))
			}
		case reflect.Slice:
			sqlType = fmt.Sprintf("ARRAY<%s>", b.dataTypeOfStruct(field.Struct))
		default:
			if _, ok := dataValue.Interface().([]byte); ok {
				sqlType = "BYTES"
			}
		}
	}
	if sqlType == "uuid" {
		sqlType = "STRING"
	}

	if sqlType == "" {
		panic(fmt.Sprintf("invalid sql type %s (%s) for commonDialect", dataValue.Type().Name(), dataValue.Kind().String()))
	}

	if strings.TrimSpace(additionalType) == "" {
		return sqlType
	}
	return fmt.Sprintf("%v %v", sqlType, additionalType)
}

func (b *bigQueryDialect) HasIndex(tableName string, indexName string) bool {
	return processor.HasIndex(b.db, tableName, indexName)
}

func (b *bigQueryDialect) HasForeignKey(tableName string, foreignKeyName string) bool {
	return processor.HasForeignKey(b.db, tableName, foreignKeyName)
}

func (b *bigQueryDialect) RemoveIndex(tableName string, indexName string) error {
	return processor.RemoveIndex(b.db, tableName, indexName)
}

func (b *bigQueryDialect) HasTable(tableName string) bool {
	return processor.HasTable(b.db, tableName)
}

func (b *bigQueryDialect) HasColumn(tableName string, columnName string) bool {
	return processor.HasColumn(b.db, tableName, columnName)
}

func (b *bigQueryDialect) ModifyColumn(tableName string, columnName string, columnType string) error {
	return processor.ModifyColumn(b.db, tableName, columnName, columnType)
}

func (b *bigQueryDialect) LimitAndOffsetSQL(limit, offset interface{}) (string, error) {

	if isValidLimitOrOffset(limit) && !isValidLimitOrOffset(offset) {
		return fmt.Sprintf(" LIMIT %d", limit), nil
	} else if isValidLimitOrOffset(limit) && isValidLimitOrOffset(offset) {
		return fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset), nil
	}
	return "", nil
}

func isValidLimitOrOffset(value interface{}) bool {
	return value != -1 && value != nil
}

func (b *bigQueryDialect) SelectFromDummyTable() string {
	return ""
}

func (bigQueryDialect) SupportLastInsertID() bool {
	return false
}

func (bigQueryDialect) LastInsertIDOutputInterstitial(tableName, columnName string, columns []string) string {
	return ""
}

func (bigQueryDialect) LastInsertIDReturningSuffix(tableName, columnName string) string {
	return ""
}

func (b *bigQueryDialect) DefaultValueStr() string {
	return ""
}

func (b *bigQueryDialect) BuildKeyName(kind, tableName string, fields ...string) string {
	return ""
}

func (b *bigQueryDialect) NormalizeIndexAndColumn(indexName, columnName string) (string, string) {
	return indexName, columnName
}

func (b *bigQueryDialect) CurrentDatabase() string {
	return ""
}

func (b *bigQueryDialect) fieldCanAutoIncrement(field *gorm.StructField) bool {
	if value, ok := field.TagSettingsGet("AUTO_INCREMENT"); ok {
		return strings.ToLower(value) != "false"
	}
	return field.IsPrimaryKey
}

func (b *bigQueryDialect) dataTypeOfStruct(field reflect.StructField) string {

	fieldType := field.Type

	var fieldDefinitions []string

	switch fieldType.Kind() {
	case reflect.Slice, reflect.Array:
		fieldType = fieldType.Elem()
	}

	for index := 0; index < fieldType.NumField(); index++ {
		subField := fieldType.Field(index)

		fieldDefinitions = append(fieldDefinitions, fmt.Sprintf("%s %s", subField.Name, b.DataTypeOf(&gorm.StructField{
			Name:        subField.Name,
			Names:       []string{subField.Name},
			Tag:         subField.Tag,
			TagSettings: parseTagSetting(subField.Tag),
			Struct:      subField,
		})))
	}

	return fmt.Sprintf("STRUCT<%s>", strings.Join(fieldDefinitions, ", "))
}

func parseTagSetting(tags reflect.StructTag) map[string]string {
	setting := map[string]string{}
	for _, str := range []string{tags.Get("sql"), tags.Get("gorm")} {
		if str == "" {
			continue
		}
		tags := strings.Split(str, ";")
		for _, value := range tags {
			v := strings.Split(value, ":")
			k := strings.TrimSpace(strings.ToUpper(v[0]))
			if len(v) >= 2 {
				setting[k] = strings.Join(v[1:], ":")
			} else {
				setting[k] = k
			}
		}
	}
	return setting
}

var parseFieldStructForDialect = func(field *gorm.StructField, dialect gorm.Dialect) (fieldValue reflect.Value, sqlType string, size int, additionalType string) {
	// Get redirected field type
	var (
		reflectType = field.Struct.Type
		dataType, _ = field.TagSettingsGet("TYPE")
	)

	for reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}

	// Get redirected field value
	fieldValue = reflect.Indirect(reflect.New(reflectType))

	if gormDataType, ok := fieldValue.Interface().(interface {
		GormDataType(gorm.Dialect) string
	}); ok {
		dataType = gormDataType.GormDataType(dialect)
	}

	// Default Size
	if num, ok := field.TagSettingsGet("SIZE"); ok {
		size, _ = strconv.Atoi(num)
	} else {
		size = 255
	}

	// Default type from tag setting
	notNull, _ := field.TagSettingsGet("NOT NULL")
	unique, _ := field.TagSettingsGet("UNIQUE")
	additionalType = notNull + " " + unique
	if value, ok := field.TagSettingsGet("DEFAULT"); ok {
		additionalType = additionalType + " DEFAULT " + value
	}

	if value, ok := field.TagSettingsGet("COMMENT"); ok {
		additionalType = additionalType + " COMMENT " + value
	}

	return fieldValue, dataType, size, strings.TrimSpace(additionalType)
}
