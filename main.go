package main

import (
	_ "github.com/guypeled76/go-bigquery-driver/gorm/dialect"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"log"
)

type RunTestProject struct {
	Name string `gorm:"column:Name"`
}

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	db, err := gorm.Open("bigquery", "bigquery://unity-rd-perf-test-data-prd/location/perf_test_results")
	if err != nil {
		log.Fatal(err)
	}

	var projects []RunTestProject

	if db.HasTable(projects) {
		log.Println("verified has table")
	}

	if db.HasTable(projects) {
		log.Println("verified has table")
	}

	err = db.Not("Name", []string{"", "2D"}).Limit(2).Find(&projects).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, project := range projects {
		log.Println(project.Name)
	}

	err = db.Not("Name", []string{"", "2D"}).Limit(2).Offset(3).Find(&projects).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, project := range projects {
		log.Println(project.Name)
	}

	err = db.Not("Name", []string{"", "2D"}).Find(&projects).Error
	if err != nil {
		log.Fatal(err)
	}

	for _, project := range projects {
		log.Println(project.Name)
	}

	defer db.Close()
	// Do Something with the DB

}
