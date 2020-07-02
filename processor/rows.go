package processor

import (
	"database/sql/driver"
	"io"
)

func createValueRows(value interface{}) driver.Rows {
	return &queryResults{
		values:  []interface{}{value},
		columns: []string{"Value"},
	}
}

type queryResults struct {
	values  []interface{}
	columns []string
}

func (results *queryResults) Columns() []string {
	return results.columns
}

func (results *queryResults) Close() error {
	return nil
}

func (results *queryResults) Next(dest []driver.Value) error {
	if results.values != nil {
		var length = len(results.values)
		for i := range dest {
			if i < length {
				dest[i] = results.values[i]
			}
		}
		results.values = nil
		return nil
	}

	return io.EOF
}
