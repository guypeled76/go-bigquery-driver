package processor

import (
	"cloud.google.com/go/bigquery"
	"database/sql/driver"
)

type processorHandler func(*bigquery.Dataset, []driver.Value) (driver.Rows, error)

var processors = map[string]processorHandler{
	hasTable: func(dataset *bigquery.Dataset, args []driver.Value) (driver.Rows, error) {
		return createValueRows(dataset.Table(args[0].(string)) != nil), nil
	},
}
