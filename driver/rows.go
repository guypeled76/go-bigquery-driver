package driver

import (
	"cloud.google.com/go/bigquery"
	"database/sql/driver"
)

type bigQueryRows struct {
	rowIterator *bigquery.RowIterator
}

func (rows bigQueryRows) Columns() []string {
	var columns []string
	for _, column := range rows.rowIterator.Schema {
		columns = append(columns, column.Name)
	}
	return columns
}

func (rows bigQueryRows) Close() error {
	return nil
}

func (rows bigQueryRows) Next(dest []driver.Value) error {
	values, err := rows.next()
	if err != nil {
		return err
	}

	dest[0] = values

	return nil
}

func (rows bigQueryRows) next() ([]bigquery.Value, error) {
	var values []bigquery.Value
	err := rows.rowIterator.Next(&values)
	if err != nil {
		return nil, err
	}

	return values, nil
}
