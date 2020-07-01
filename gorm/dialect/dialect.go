package dialect

import (
	"errors"
	"fmt"
	_ "github.com/guypeled76/go-bigquery-driver/driver"
	"github.com/guypeled76/go-bigquery-driver/processor"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"strings"
	"time"
)

const Name = "bigquery"

type bigQueryDialect struct {
	db gorm.SQLCommon
}

func (b bigQueryDialect) GetName() string {
	return Name
}

func (b bigQueryDialect) SetDB(db gorm.SQLCommon) {
	b.db = db
}

func (b bigQueryDialect) BindVar(i int) string {
	unsupportedPanic("BindVar")
	return ""
}

func (b bigQueryDialect) Quote(key string) string {
	return fmt.Sprintf("`%s`", key)
}

func (b bigQueryDialect) DataTypeOf(field *gorm.StructField) string {
	var dataValue, sqlType, _, additionalType = gorm.ParseFieldStructForDialect(field, b)
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
			}
		case reflect.Array:
			if _, ok := dataValue.Interface().(uuid.UUID); ok {
				sqlType = "STRING"
			}
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

func (b bigQueryDialect) HasIndex(tableName string, indexName string) bool {
	return processor.HasIndex(b.db, tableName, indexName)
}

func (b bigQueryDialect) HasForeignKey(tableName string, foreignKeyName string) bool {
	return processor.HasForeignKey(b.db, tableName, foreignKeyName)
}

func (b bigQueryDialect) RemoveIndex(tableName string, indexName string) error {
	return unsupportedError("RemoveIndex")
}

func (b bigQueryDialect) HasTable(tableName string) bool {
	return processor.HasTable(b.db, tableName)
}

func (b bigQueryDialect) HasColumn(tableName string, columnName string) bool {
	return processor.HasColumn(b.db, tableName, columnName)
}

func (b bigQueryDialect) ModifyColumn(tableName string, columnName string, columnType string) error {
	return processor.ModifyColumn(b.db, tableName, columnName, columnType)
}

func (b bigQueryDialect) LimitAndOffsetSQL(limit, offset interface{}) (string, error) {
	return "", nil
}

func (b bigQueryDialect) SelectFromDummyTable() string {
	panic("implement me")
}

func (b bigQueryDialect) LastInsertIDOutputInterstitial(tableName, columnName string, columns []string) string {
	panic("implement me")
}

func (b bigQueryDialect) LastInsertIDReturningSuffix(tableName, columnName string) string {
	panic("implement me")
}

func (b bigQueryDialect) DefaultValueStr() string {
	panic("implement me")
}

func (b bigQueryDialect) BuildKeyName(kind, tableName string, fields ...string) string {
	panic("implement me")
}

func (b bigQueryDialect) NormalizeIndexAndColumn(indexName, columnName string) (string, string) {
	panic("implement me")
}

func (b bigQueryDialect) CurrentDatabase() string {
	panic("implement me")
}

func (b *bigQueryDialect) fieldCanAutoIncrement(field *gorm.StructField) bool {
	if value, ok := field.TagSettingsGet("AUTO_INCREMENT"); ok {
		return strings.ToLower(value) != "false"
	}
	return field.IsPrimaryKey
}

func unsupportedError(feature string) error {
	return errors.New(unsupportedMessage(feature))
}

func unsupportedPanic(feature string) {
	panic(unsupportedMessage(feature))
}

func unsupportedMessage(feature string) string {
	return fmt.Sprintf("BigQuery GORM dialect doesn't support '%s'", feature)
}

func uninitializedError(feature string) error {
	return errors.New(uninitializedMessage(feature))
}

func uninitializedPanic(feature string) {
	panic(uninitializedMessage(feature))
}

func uninitializedMessage(feature string) string {
	return fmt.Sprintf("BigQuery GORM dialect '%s' was called before initializing db", feature)
}
