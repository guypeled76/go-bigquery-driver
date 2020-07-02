package processor

import (
	"cloud.google.com/go/bigquery"
	"database/sql/driver"
)

type Provider interface {
	GetDataset() *bigquery.Dataset
}

func Exec(provider Provider, query string, args []driver.Value) (driver.Result, error) {
	return nil, nil
}

func Query(provider Provider, query string, args []driver.Value) (driver.Rows, error) {

	handler := processors[query]
	if handler != nil {
		return handler(provider, args)
	}

	return nil, nil
}
