package driver

import (
	"cloud.google.com/go/bigquery"
	"errors"
	"io"
)

type bigQuerySource interface {
	GetSchema() bigQuerySchema
	Next() ([]bigquery.Value, error)
}

type bigQueryRowIteratorSource struct {
	iterator *bigquery.RowIterator
}

func (source *bigQueryRowIteratorSource) GetSchema() bigQuerySchema {
	return createBigQuerySchema(source.iterator.Schema)
}

func (source *bigQueryRowIteratorSource) Next() ([]bigquery.Value, error) {
	var values []bigquery.Value
	err := source.iterator.Next(&values)
	return values, err
}

func createSourceFromRowIterator(rowIterator *bigquery.RowIterator) bigQuerySource {
	return &bigQueryRowIteratorSource{
		rowIterator,
	}
}

type bigQueryColumnSource struct {
	schema   bigQuerySchema
	rows     []bigquery.Value
	position int
}

func (source *bigQueryColumnSource) GetSchema() bigQuerySchema {
	return source.schema
}

func (source *bigQueryColumnSource) Next() ([]bigquery.Value, error) {
	if source.position >= len(source.rows) {
		return nil, io.EOF
	}
	values, ok := source.rows[source.position].([]bigquery.Value)
	if !ok {
		return nil, errors.New("failed to get row from column source")
	}
	source.position++
	return values, nil
}

func createSourceFromColumn(schema bigQuerySchema, rows []bigquery.Value) bigQuerySource {
	return &bigQueryColumnSource{
		schema:   schema,
		rows:     rows,
		position: 0,
	}
}
