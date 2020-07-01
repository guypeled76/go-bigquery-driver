package driver

type bigQueryResult struct {
}

func (result bigQueryResult) LastInsertId() (int64, error) {
	panic("implement me")
}

func (result bigQueryResult) RowsAffected() (int64, error) {
	panic("implement me")
}
