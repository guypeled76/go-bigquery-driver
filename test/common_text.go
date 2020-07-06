package test

import (
	_ "github.com/guypeled76/go-bigquery-driver/gorm/dialect"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"log"
)

type GormTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *GormTestSuite) SetupSuite() {

	logrus.SetLevel(logrus.DebugLevel)

	var err error
	suite.db, err = gorm.Open("bigquery", "bigquery://go-bigquery-driver/playground")
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *GormTestSuite) TearDownSuite() {
	suite.db.Close()
}
