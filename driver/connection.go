package driver

import (
	"cloud.google.com/go/bigquery"
	"context"
	"database/sql/driver"
	"fmt"
)

type bigQueryConnection struct {
	ctx     context.Context
	client  *bigquery.Client
	config  bigQueryConfig
	closed  bool
	bad     bool
	dataset *bigquery.Dataset
}

func (connection *bigQueryConnection) GetDataset() *bigquery.Dataset {
	if connection.dataset != nil {
		return connection.dataset
	}
	connection.dataset = connection.client.Dataset(connection.config.dataSet)
	return connection.dataset
}

func (connection *bigQueryConnection) GetContext() context.Context {
	return connection.ctx
}

func (connection *bigQueryConnection) Ping(ctx context.Context) error {

	dataset := connection.GetDataset()
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
	if connection.closed {
		return nil
	}
	if connection.bad {
		return driver.ErrBadConn
	}
	connection.closed = true
	return connection.client.Close()
}

func (connection *bigQueryConnection) Begin() (driver.Tx, error) {
	var transaction = &bigQueryTransaction{connection}

	return transaction, nil
}

func (connection *bigQueryConnection) query(query string) (*bigquery.Query, error) {
	return connection.client.Query(query), nil
}
