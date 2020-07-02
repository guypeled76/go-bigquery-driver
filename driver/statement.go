package driver

import (
	"cloud.google.com/go/bigquery"
	"context"
	"database/sql/driver"
	"github.com/guypeled76/go-bigquery-driver/processor"
	"github.com/sirupsen/logrus"
)

type bigQueryStatement struct {
	connection *bigQueryConnection
	query      string
}

func (statement bigQueryStatement) Close() error {
	return nil
}

func (statement bigQueryStatement) NumInput() int {
	return 0
}

func (statement bigQueryStatement) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return statement.Exec(convertParameters(args))
}

func (statement bigQueryStatement) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return statement.Query(convertParameters(args))
}

func (statement bigQueryStatement) Exec(args []driver.Value) (driver.Result, error) {
	result, err := processor.Exec(statement.query, args)
	if err != nil || result != nil {
		return result, err
	}

	query, err := statement.buildQuery(args)
	if err != nil {
		return nil, err
	}

	rowIterator, err := query.Read(context.Background())
	if err != nil {
		return nil, err
	}

	return &bigQueryResult{rowIterator}, nil
}

func (statement bigQueryStatement) Query(args []driver.Value) (driver.Rows, error) {
	rows, err := processor.Query(statement.query, args)
	if err != nil || rows != nil {
		return rows, err
	}

	query, err := statement.buildQuery(args)
	if err != nil {
		return nil, err
	}

	rowIterator, err := query.Read(context.Background())
	if err != nil {
		return nil, err
	}

	return &bigQueryRows{rowIterator: rowIterator}, nil
}

func (statement bigQueryStatement) buildQuery(args []driver.Value) (*bigquery.Query, error) {

	logrus.Debugf("query:%s", statement.query)

	query, err := statement.connection.query(statement.query)
	if err != nil {
		return nil, err
	}
	query.DefaultDatasetID = statement.connection.config.dataSet
	query.Parameters, err = statement.buildParameters(args)
	if err != nil {
		return nil, err
	}

	return query, err
}

func (statement bigQueryStatement) buildParameters(args []driver.Value) ([]bigquery.QueryParameter, error) {
	if args == nil {
		return nil, nil
	}

	var parameters []bigquery.QueryParameter
	for _, arg := range args {
		parameters = buildParameter(arg, parameters)
	}
	return parameters, nil
}

func buildParameter(arg driver.Value, parameters []bigquery.QueryParameter) []bigquery.QueryParameter {
	namedValue, ok := arg.(driver.NamedValue)
	if ok && namedValue.Name != "" {

		logrus.Debugf("-param:%s=%s", namedValue.Name, namedValue.Value)

		return append(parameters, bigquery.QueryParameter{
			Name:  namedValue.Name,
			Value: namedValue.Value,
		})
	}

	logrus.Debugf("-param:%s", arg)

	return append(parameters, bigquery.QueryParameter{
		Value: arg,
	})
}

func convertParameters(args []driver.NamedValue) []driver.Value {
	var values []driver.Value
	if args != nil {
		for _, arg := range args {
			values = append(values, arg)
		}
	}
	return values
}
