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

func (suite *MetadataTestSuit) Test_CreateSimpleTable() {
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

func (suite *MetadataTestSuit) Test_CreateComplexTable() {
	suite.db.DropTableIfExists(&ComplexRecord{})
	assert.False(suite.T(), suite.db.HasTable(&ComplexRecord{}))
	suite.db.AutoMigrate(&ComplexRecord{})
	assert.True(suite.T(), suite.db.HasTable(&ComplexRecord{}))
}

func (suite *MetadataTestSuit) Test_SelectComplexTable() {
	var records []ComplexRecord
	suite.db.Find(&records)

	if len(records) > 0 {

	}
}

func (suite *MetadataTestSuit) Test_SelectArrayTable() {
	var records []ArrayRecord
	suite.db.DropTableIfExists(&ArrayRecord{})
	assert.False(suite.T(), suite.db.HasTable(&ArrayRecord{}))
	suite.db.AutoMigrate(&ArrayRecord{})
	assert.True(suite.T(), suite.db.HasTable(&ArrayRecord{}))
	suite.db.Create(&ArrayRecord{Name: "test", Records: ArrayRecordRecord{{Name: "dd", Age: 1}, {Name: "dd1", Age: 1}}})
	suite.db.Create(&ArrayRecord{Name: "test2", Records: ArrayRecordRecord{{Name: "dd2", Age: 444}, {Name: "dd3", Age: 1}}})
	suite.db.Order("Name").Find(&records)
	assert.Equal(suite.T(), 2, len(records), "we should have two records")
	if len(records) == 2 {
		assert.Equal(suite.T(), 444, records[1].Records[0].Age)
	}
}
