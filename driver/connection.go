package driver

import (
	"cloud.google.com/go/bigquery"
	"context"
	"database/sql/driver"
	"fmt"
)

type bigQueryConnection struct {
	client *bigquery.Client
	config bigQueryConfig
}

func (connection *bigQueryConnection) Ping(ctx context.Context) error {

	dataset := connection.client.Dataset(connection.config.dataSet)
	if dataset == nil {
		return fmt.Errorf("faild to ping using '%s' dataset", connection.config.dataSet)
	}

	_, err := dataset.Metadata(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (connection *bigQueryConnection) Query(query string, args []driver.Value) (driver.Rows, error) {
	statement, err := connection.Prepare(query)
	if err != nil {
		return nil, nil
	}

	return statement.Query(args)
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

func (connection *bigQueryConnection) query(query string) (*bigquery.Query, error) {
	return connection.client.Query(query), nil
}
