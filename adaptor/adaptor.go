package adaptor

import (
	"context"
	"database/sql/driver"
)

const (
	RecordType = "RECORD"
	ArrayType  = "ARRAY"

	RerouteQuery = "SELECT ?"
)

var adaptorCtxKey = struct{ value string }{"adaptorCxtKey"}

// SchemaAdaptor adapts row column results to fit the "demand"
type SchemaAdaptor interface {

	// GetColumnAdaptor gets a specific column adapter that makes the column fit the "demand"
	GetColumnAdaptor(name string) SchemaColumnAdaptor
}

// SchemaColumnAdaptor adapts column results to fit the "demand"
type SchemaColumnAdaptor interface {

	// AdaptValue gets a specific column value that fit the "demand"
	AdaptValue(value driver.Value) driver.Value

	// GetSchemaAdaptor handles a nested schema that needs to be also adapted to the "demand"
	GetSchemaAdaptor() SchemaAdaptor
}

func SetSchemaAdaptor(ctx context.Context, adaptorSchema SchemaAdaptor) context.Context {
	if ctx != nil {
		return context.WithValue(ctx, adaptorCtxKey, adaptorSchema)
	}
	return nil
}

func GetSchemaAdaptor(ctx context.Context) SchemaAdaptor {
	if ctx == nil {
		return nil
	}

	value := ctx.Value(adaptorCtxKey)
	if value == nil {
		return nil
	}
	return value.(SchemaAdaptor)
}
