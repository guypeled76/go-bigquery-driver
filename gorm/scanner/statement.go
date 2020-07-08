package scanner

import (
	"database/sql/driver"
	"errors"
)

type scannerStatement struct {
}

func (scannerStatement) CheckNamedValue(*driver.NamedValue) error {
	// TODO: Revise in the future
	return nil
}

func (s scannerStatement) Close() error {
	return nil
}

func (s scannerStatement) NumInput() int {
	return 1
}

func (s scannerStatement) Exec(args []driver.Value) (driver.Result, error) {
	return nil, errors.New("execution is not supported")
}

func (s scannerStatement) Query(args []driver.Value) (driver.Rows, error) {
	if len(args) < 1 {
		return nil, errors.New("scanner arguments should have an argument with rows")
	}

	rows, ok := args[0].(driver.Rows)
	if !ok {
		return nil, errors.New("scanner arguments should have an argument with rows")
	}

	return rows, nil
}
