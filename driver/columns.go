package driver

import (
	"cloud.google.com/go/bigquery"
	"database/sql/driver"
)

type bigQuerySchema interface {
	ColumnNames() []string
	ConvertColumnValue(index int, value bigquery.Value) driver.Value
}

type bigQueryColumns struct {
	names   []string
	columns []bigQueryColumn
}

func (columns bigQueryColumns) ConvertColumnValue(index int, value bigquery.Value) driver.Value {
	if index > -1 && len(columns.columns) > index {
		column := columns.columns[index]
		return column.ConvertValue(value)
	}

	return value
}

func (columns bigQueryColumns) ColumnNames() []string {
	return columns.names
}

type bigQueryColumn struct {
	Name   string
	Schema bigquery.Schema
}

func (column bigQueryColumn) ConvertValue(value bigquery.Value) driver.Value {

	if len(column.Schema) == 0 {
		return value
	}

	values, ok := value.([]bigquery.Value)
	if ok {

		if len(values) > 0 {
			if _, isRows := values[0].([]bigquery.Value); !isRows {
				values = []bigquery.Value{values}
			}
		}

		schema := createBigQuerySchema(column.Schema)

		return &bigQueryRows{
			source: createSourceFromColumn(schema, values),
		}
	}

	return value
}

func createBigQuerySchema(schema bigquery.Schema) bigQuerySchema {
	var names []string
	var columns []bigQueryColumn
	for _, column := range schema {
		names = append(names, column.Name)
		columns = append(columns, bigQueryColumn{
			Name:   column.Name,
			Schema: column.Schema,
		})
	}
	return &bigQueryColumns{
		names,
		columns,
	}
}
