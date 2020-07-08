package scanner

import "database/sql/driver"

type scannerDriver struct {
}

func (s scannerDriver) Open(name string) (driver.Conn, error) {
	return scannerConnection{}, nil
}
