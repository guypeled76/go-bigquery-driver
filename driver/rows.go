package driver

import (
	"cloud.google.com/go/bigquery"
	"database/sql/driver"
	"google.golang.org/api/iterator"
	"io"
)

type bigQueryRows struct {
	source bigQuerySource
	schema bigQuerySchema
}

func (rows *bigQueryRows) ensureSchema() {
	if rows.schema == nil {
		rows.schema = rows.source.GetSchema()
	}
}

func (rows *bigQueryRows) Columns() []string {
	rows.ensureSchema()
	return rows.schema.ColumnNames()
}

func (rows *bigQueryRows) Close() error {
	return nil
}

func (rows *bigQueryRows) Next(dest []driver.Value) error {

	rows.ensureSchema()

	var values []bigquery.Value
	err := rows.source.Next(&values)
	if err == iterator.Done {
		return io.EOF
	}

	if err != nil {
		return err
	}

	var length = len(values)
	for i := range dest {
		if i < length {
			dest[i] = rows.schema.ConvertColumnValue(i, values[i])
		}
	}

	return nil
}
