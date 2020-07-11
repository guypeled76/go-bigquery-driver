package main

import (
	"github.com/guypeled76/go-bigquery-driver/bigquery"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
)

type ComplexSubRecord struct {
	Name string `gorm:"column:Name"`
	Age  int    `gorm:"column:Age"`
}

type ArrayRecord struct {
	Name    string             `gorm:"column:Name"`
	Records []ComplexSubRecord `gorm:"column:Records;type:ARRAY"`
}

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	db, err := gorm.Open(bigquery.Open("bigquery://go-bigquery-driver/playground"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var records []ArrayRecord
	db.Migrator().DropTable(&ArrayRecord{})
	db.AutoMigrate(&ArrayRecord{})
	db.Create(&ArrayRecord{Name: "test", Records: []ComplexSubRecord{{Name: "dd", Age: 1}, {Name: "dd1", Age: 1}}})
	db.Create(&ArrayRecord{Name: "test2", Records: []ComplexSubRecord{{Name: "dd2", Age: 444}, {Name: "dd3", Age: 1}}})
	db.Order("Name").Find(&records)

	for _, record := range records {
		for _, subRecord := range record.Records {
			log.Printf("%s=%s", subRecord.Name, subRecord.Name)
		}
	}

}
