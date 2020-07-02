package processor

import (
	"cloud.google.com/go/bigquery"
	"database/sql/driver"
	"github.com/sirupsen/logrus"
)

type DatasetProvider interface {
	GetDataset() *bigquery.Dataset
}

func Exec(provider DatasetProvider, query string, args []driver.Value) (driver.Result, error) {
	return nil, nil
}

func Query(provider DatasetProvider, query string, args []driver.Value) (driver.Rows, error) {

	handler := processors[query]
	if handler != nil {

		logrus.Debugf("processor:%s", query)

		return handler(provider.GetDataset(), args)
	}

	return nil, nil
}
