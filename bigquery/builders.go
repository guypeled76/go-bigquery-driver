package bigquery

import (
	"database/sql"
	"database/sql/driver"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"reflect"
)

func initializeBuilders(db *gorm.DB) {
	db.ClauseBuilders["VALUES"] = valuesBuilder
}

func valuesBuilder(c clause.Clause, builder clause.Builder) {

	if c.Expression == nil {
		return
	}

	values, ok := c.Expression.(clause.Values)
	if !ok {
		return
	}
	log.Print("f")

	if len(values.Columns) > 0 {
		builder.WriteByte('(')
		for idx, column := range values.Columns {
			if idx > 0 {
				builder.WriteByte(',')
			}
			builder.WriteQuoted(column)
		}
		builder.WriteByte(')')

		builder.WriteString(" VALUES ")

		for idx, value := range values.Values {
			if idx > 0 {
				builder.WriteByte(',')
			}

			builder.WriteByte('(')
			valuesBuilderBuildValues(builder, value)
			builder.WriteByte(')')
		}
	} else {
		builder.WriteString("DEFAULT VALUES")
	}
}

func valuesBuilderBuildValues(builder clause.Builder, vars []interface{}) {
	for idx, v := range vars {
		if idx > 0 {
			builder.WriteByte(',')
		}

		switch v := v.(type) {
		case sql.NamedArg, clause.Column, clause.Table, clause.Expr, driver.Valuer, []byte, []interface{}, *gorm.DB:
			builder.AddVar(builder, v)
		default:
			switch rv := reflect.ValueOf(v); rv.Kind() {
			case reflect.Slice, reflect.Array:
				if rv.Len() == 0 {
					builder.WriteString("[]")
				} else {
					builder.WriteByte('[')
					for i := 0; i < rv.Len(); i++ {
						if i > 0 {
							builder.WriteByte(',')
						}
						builder.AddVar(builder, rv.Index(i).Interface())
					}
					builder.WriteByte(']')
				}
			default:
				builder.AddVar(builder, v)
			}
		}
	}
}
