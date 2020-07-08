package scanner

import (
	"context"
	"database/sql/driver"
)

type scannerConnection struct {
}

func (scannerConnection) Prepare(query string) (driver.Stmt, error) {
	return &scannerStatement{}, nil
}

func (scannerConnection) Close() error {
	return nil
}

func (scannerConnection) Begin() (driver.Tx, error) {
	return nil, nil
}

func (scannerConnection) Ping(ctx context.Context) error {
	return nil
}

func (scannerConnection) CheckNamedValue(*driver.NamedValue) error {
	// TODO: Revise in the future
	return nil
}
