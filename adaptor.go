package bigquery

import (
	"database/sql/driver"
	"github.com/go-gorm/bigquery/adaptor"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"reflect"
)

type bigQuerySchemaAdaptor struct {
	schema *schema.Schema
	db     *gorm.DB
}

func (schemaAdaptor *bigQuerySchemaAdaptor) GetColumnAdaptor(name string) adaptor.SchemaColumnAdaptor {

	if schema := schemaAdaptor.schema; schema != nil {

		field := schema.FieldsByDBName[name]

		if field == nil {
			return nil
		}

		switch field.DataType {
		case adaptor.RecordType, adaptor.ArrayType:
			return &bigQueryColumnAdaptor{field: field, db: schemaAdaptor.db}
		}
	}

	return nil
}

type bigQueryColumnAdaptor struct {
	field *schema.Field
	db    *gorm.DB
}

func (columnAdaptor bigQueryColumnAdaptor) AdaptValue(value driver.Value) driver.Value {
	instance := reflect.New(columnAdaptor.field.IndirectFieldType).Interface()
	columnAdaptor.db.Raw(adaptor.RerouteQuery, value).Scan(instance)
	return instance
}

func (columnAdaptor bigQueryColumnAdaptor) GetSchemaAdaptor() adaptor.SchemaAdaptor {
	schema := columnAdaptor.field.Schema

	if schema == nil {
		return nil
	}
	return &bigQuerySchemaAdaptor{
		schema: schema,
		db:     columnAdaptor.db,
	}
}
