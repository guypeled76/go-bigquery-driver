package driver

import (
	"cloud.google.com/go/bigquery"
	"database/sql/driver"
)

type bigQueryConnection struct {
	client *bigquery.Client
	config bigQueryConfig
}

func (connection *bigQueryConnection) Prepare(query string) (driver.Stmt, error) {
	var statement = &bigQueryStatement{connection, query}

	return statement, nil
}

func (connection *bigQueryConnection) Close() error {

	return nil
}

func (connection *bigQueryConnection) Begin() (driver.Tx, error) {
	var transaction = &bigQueryTransaction{connection}

	return transaction, nil
}
