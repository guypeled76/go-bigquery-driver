package processor

import (
	"context"
	"database/sql/driver"
	"errors"
)

type processorHandler func(Provider, []driver.Value) (driver.Rows, error)

var processors = map[string]processorHandler{
	hasTable:  processHasTable,
	hasColumn: processHasColumn,
}

func processHasTable(provider Provider, args []driver.Value) (driver.Rows, error) {
	_, err := provider.GetDataset().Table(args[0].(string)).Metadata(context.Background())
	return rowsFromValue(err == nil), nil
}

func processHasColumn(provider Provider, args []driver.Value) (driver.Rows, error) {
	if len(args) < 2 {
		return nil, errors.New("has column needs two arguments")
	}
	hasColumnFlag, err := evaluateHasColumn(provider, args[0].(string), args[1].(string))
	return rowsFromValue(hasColumnFlag), err
}

func evaluateHasColumn(provider Provider, tableName, columnName string) (bool, error) {
	metadata, err := provider.GetDataset().Table(tableName).Metadata(context.Background())
	if err != nil {
		return false, err
	}

	for _, field := range metadata.Schema {
		if field.Name == columnName {
			return true, nil
		}
	}

	return false, nil
}
