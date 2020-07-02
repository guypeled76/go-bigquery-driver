package processor

import (
	"database/sql/driver"
)

type processorHandler func(Provider, []driver.Value) (driver.Rows, error)

var processors = map[string]processorHandler{
	hasTable: processHasTable,
}

func processHasTable(provider Provider, args []driver.Value) (driver.Rows, error) {
	return createValueRows(provider.GetDataset().Table(args[0].(string)) != nil), nil
}
