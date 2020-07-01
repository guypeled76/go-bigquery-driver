package driver

import (
	"database/sql/driver"
	"github.com/guypeled76/go-bigquery-driver/processor"
)

type bigQueryStatement struct {
	connection *bigQueryConnection
	query      string
}

func (statement bigQueryStatement) Close() error {
	return nil
}

func (statement bigQueryStatement) NumInput() int {
	return 0
}

func (statement bigQueryStatement) Exec(args []driver.Value) (driver.Result, error) {
	result, err := processor.Exec(statement.query, args)
	if err != nil || result != nil {
		return result, err
	}

	result = &bigQueryResult{}

	return result, nil
}

func (statement bigQueryStatement) Query(args []driver.Value) (driver.Rows, error) {
	rows, err := processor.Query(statement.query, args)
	if err != nil || rows != nil {
		return rows, err
	}

	rows = &bigQueryRows{}

	return rows, nil
}
