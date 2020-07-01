package processor

import "github.com/jinzhu/gorm"

func HasTable(db gorm.SQLCommon, tableName string) bool {
	return queryToBool(db, hasTable, tableName)
}

func HasIndex(db gorm.SQLCommon, tableName string, indexName string) bool {
	return queryToBool(db, hasIndex, tableName, indexName)
}

func HasForeignKey(db gorm.SQLCommon, tableName string, foreignKeyName string) bool {
	return queryToBool(db, hasForeignKey, tableName, foreignKeyName)
}

func HasColumn(db gorm.SQLCommon, tableName string, columnName string) bool {
	return queryToBool(db, hasColumn, tableName, columnName)
}

func ModifyColumn(db gorm.SQLCommon, tableName string, columnName string, columnType string) error {
	return execute(db, modifyColumn, tableName, columnName, columnType)
}
