package driver

import (
	"cloud.google.com/go/bigquery"
	"context"
	"database/sql/driver"
	"fmt"
	"strings"
)

type bigQueryDriver struct {
}

type bigQueryConfig struct {
	projectID string
	location  string
	dataSet   string
}

func (b bigQueryDriver) Open(uri string) (driver.Conn, error) {

	if uri == "scanner" {
		return &scannerConnection{}, nil
	}

	config, err := configFromUri(uri)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, config.projectID)
	if err != nil {
		return nil, err
	}

	return &bigQueryConnection{
		ctx:    ctx,
		client: client,
		config: *config,
	}, nil
}

func configFromUri(uri string) (*bigQueryConfig, error) {

	if !strings.HasPrefix(uri, "bigquery://") {
		return nil, fmt.Errorf("invalid prefix, expected bigquery:// got: %s", uri)
	}

	uri = strings.ToLower(uri)
	path := strings.TrimPrefix(uri, "bigquery://")
	fields := strings.Split(path, "/")

	if len(fields) == 3 {
		return &bigQueryConfig{
			projectID: fields[0],
			location:  fields[1],
			dataSet:   fields[2],
		}, nil
	}

	if len(fields) == 2 {
		return &bigQueryConfig{
			projectID: fields[0],
			location:  "",
			dataSet:   fields[1],
		}, nil
	}

	return nil, fmt.Errorf("invalid connection string : %s", uri)

}
