package driver

import "errors"

type bigQueryResult struct {
	statement bigQueryStatement
}

func (result bigQueryResult) LastInsertId() (int64, error) {
	return 0, errors.New("LastInsertId is not supported")
}

func (result bigQueryResult) RowsAffected() (int64, error) {
	panic("implement me")
}
