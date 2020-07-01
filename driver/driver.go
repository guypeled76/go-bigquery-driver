package driver

import "database/sql/driver"

type bigQueryDriver struct {
}

func (b bigQueryDriver) Open(name string) (driver.Conn, error) {
	var conn = &bigQueryConnection{}

	return conn, nil
}
