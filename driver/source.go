package driver

import "cloud.google.com/go/bigquery"

type bigQuerySource interface {
	GetSchema() bigQuerySchema
	Next(dest interface{}) error
}

type bigQueryRowIteratorSource struct {
	iterator *bigquery.RowIterator
}

func (source *bigQueryRowIteratorSource) GetSchema() bigQuerySchema {
	return createBigQueryColumns(source.iterator.Schema)
}

func (source *bigQueryRowIteratorSource) Next(dest interface{}) error {
	return source.iterator.Next(dest)
}

func createSourceFromRowIterator(rowIterator *bigquery.RowIterator) bigQuerySource {
	return &bigQueryRowIteratorSource{
		rowIterator,
	}
}
