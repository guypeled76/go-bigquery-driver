package processor

import (
	"database/sql/driver"
)

func Exec(query string, args []driver.Value) (driver.Result, error) {
	return nil, nil
}

func Query(query string, args []driver.Value) (driver.Rows, error) {
	return nil, nil
}
