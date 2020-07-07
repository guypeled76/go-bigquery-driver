package scanner

import (
	"context"
	"database/sql/driver"
)

type scannerConnector struct {
}

func (scannerConnector) Prepare(query string) (driver.Stmt, error) {
	return &scannerStatement{}, nil
}

func (scannerConnector) Close() error {
	return nil
}

func (scannerConnector) Begin() (driver.Tx, error) {
	return nil, nil
}

func (scannerConnector) Ping(ctx context.Context) error {
	return nil
}
