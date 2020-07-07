package scanner

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type scannerDialect struct {
	db gorm.SQLCommon
}

func (s scannerDialect) GetName() string {
	return driverKey
}

func (s *scannerDialect) SetDB(db gorm.SQLCommon) {
	s.db = db
}

func (s scannerDialect) BindVar(i int) string {
	return "?"
}

func (s scannerDialect) Quote(key string) string {
	return fmt.Sprintf("`%s`", key)
}

func (s scannerDialect) DataTypeOf(field *gorm.StructField) string {
	panic("implement me")
}

func (s scannerDialect) HasIndex(tableName string, indexName string) bool {
	panic("implement me")
}

func (s scannerDialect) HasForeignKey(tableName string, foreignKeyName string) bool {
	panic("implement me")
}

func (s scannerDialect) RemoveIndex(tableName string, indexName string) error {
	panic("implement me")
}

func (s scannerDialect) HasTable(tableName string) bool {
	panic("implement me")
}

func (s scannerDialect) HasColumn(tableName string, columnName string) bool {
	panic("implement me")
}

func (s scannerDialect) ModifyColumn(tableName string, columnName string, typ string) error {
	panic("implement me")
}

func (s scannerDialect) LimitAndOffsetSQL(limit, offset interface{}) (string, error) {
	return "", nil
}

func (s scannerDialect) SelectFromDummyTable() string {
	panic("implement me")
}

func (s scannerDialect) LastInsertIDOutputInterstitial(tableName, columnName string, columns []string) string {
	panic("implement me")
}

func (s scannerDialect) LastInsertIDReturningSuffix(tableName, columnName string) string {
	panic("implement me")
}

func (s scannerDialect) DefaultValueStr() string {
	panic("implement me")
}

func (s scannerDialect) BuildKeyName(kind, tableName string, fields ...string) string {
	panic("implement me")
}

func (s scannerDialect) NormalizeIndexAndColumn(indexName, columnName string) (string, string) {
	panic("implement me")
}

func (s scannerDialect) CurrentDatabase() string {
	panic("implement me")
}
