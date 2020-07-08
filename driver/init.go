package driver

import (
	"database/sql"
)

const scannerKey = "bigquery_scanner"

func init() {
	sql.Register("bigquery", &bigQueryDriver{})
}
