package driver

import (
	"database/sql/driver"
)

type bigQueryRows struct {
}

func (rows bigQueryRows) Columns() []string {
	panic("implement me")
}

func (rows bigQueryRows) Close() error {
	panic("implement me")
}

func (rows bigQueryRows) Next(dest []driver.Value) error {
	panic("implement me")
}
