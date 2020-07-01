package processor

import "github.com/jinzhu/gorm"

func queryToBool(db gorm.SQLCommon, query string, args ...interface{}) bool {
	if db == nil {
		return false
	}

	db.Query(query, args...)

	return false
}

func execute(db gorm.SQLCommon, query string, args ...interface{}) error {
	if db == nil {
		return nil
	}

	db.Exec(query, args...)

	return nil
}
