package processor

import "github.com/jinzhu/gorm"

func queryToBool(db gorm.SQLCommon, query string, args ...interface{}) bool {
	if db == nil {
		return false
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return false
	}

	var value interface{}
	rows.Next()
	rows.Scan(&value)
	return value == true
}

func execute(db gorm.SQLCommon, query string, args ...interface{}) error {
	if db == nil {
		return nil
	}

	db.Exec(query, args...)

	return nil
}
