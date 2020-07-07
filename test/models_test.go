package test

import (
	"github.com/guypeled76/go-bigquery-driver/gorm/scanner"
)

type SimpleTestRecord struct {
	Name string `gorm:"column:Name"`
}

type ComplexRecord struct {
	Name   string           `gorm:"column:Name"`
	Record ComplexSubRecord `gorm:"column:Record"`
}

type ComplexSubRecord struct {
	Name string `gorm:"column:Name"`
	Age  int    `gorm:"column:Age"`
}

func (record *ComplexSubRecord) Scan(value interface{}) error {
	return scanner.Scan(value, record)
}
