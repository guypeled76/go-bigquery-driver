package driver

type bigQueryTransaction struct {
	connection *bigQueryConnection
}

func (transaction *bigQueryTransaction) Commit() error {

	return nil
}

func (transaction *bigQueryTransaction) Rollback() error {

	return nil
}
