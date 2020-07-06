package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MetadataTestSuit struct {
	GormTestSuite
}

func TestMetadataTestSuit(t *testing.T) {
	suite.Run(t, new(MetadataTestSuit))
}

func (suite *MetadataTestSuit) Test_HasTable() {
	assert.False(suite.T(), suite.db.HasTable("non_existing_table"))
}

func (suite *MetadataTestSuit) Test_CreateTable() {
	suite.db.DropTableIfExists(&SimpleTestRecord{})
	assert.False(suite.T(), suite.db.HasTable(&SimpleTestRecord{}))
	suite.db.AutoMigrate(&SimpleTestRecord{})
	assert.True(suite.T(), suite.db.HasTable(&SimpleTestRecord{}))
	suite.db.Create(&SimpleTestRecord{Name: "test"})

	var records []SimpleTestRecord
	suite.db.First(&records)

	assert.Equal(suite.T(), 1, len(records), "should be a records")
	if len(records) > 0 {
		assert.Equal(suite.T(), "test", records[0].Name)
	}
}
