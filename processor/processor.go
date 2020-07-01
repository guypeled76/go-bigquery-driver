package processor

import (
	"database/sql/driver"
	"github.com/jinzhu/gorm"
)

const (
	hasColumn    string = "__hasColumn($tableName,$columnName)"
	modifyColumn string = "__modifyColumn($tableName,$columnName,$columnType)"
	hasTable     string = "__hasTable($tableName)"
)

func HasTable(db gorm.SQLCommon, tableName string) bool {
	return queryToBool(db, hasTable, tableName)
}

func HasColumn(db gorm.SQLCommon, tableName string, columnName string) bool {
	return queryToBool(db, hasColumn, tableName, columnName)
}

func ModifyColumn(db gorm.SQLCommon, tableName string, columnName string, columnType string) error {
	return execute(db, modifyColumn, tableName, columnName, columnType)
}

func Exec(query string, args []driver.Value) (driver.Result, error) {
	return nil, nil
}

func Query(query string, args []driver.Value) (driver.Rows, error) {
	return nil, nil
}
