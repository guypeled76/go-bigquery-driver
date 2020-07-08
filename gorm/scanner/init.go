package scanner

import (
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func init() {

	var err error
	db, err = gorm.Open("bigquery", "scanner")
	if err != nil {
		log.Fatal(err)
	}
}
