package scanner

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"log"
)

const driverKey = "bigquery_scanner"

var db *gorm.DB

func init() {

	sql.Register(driverKey, &scannerDriver{})
	gorm.RegisterDialect(driverKey, &scannerDialect{})

	var err error
	db, err = gorm.Open(driverKey, driverKey)
	if err != nil {
		log.Fatal(err)
	}
}
