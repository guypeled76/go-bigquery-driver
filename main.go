package main

import (
	_ "github.com/guypeled76/go-bigquery-driver/bigquery/dialect"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type RunTestProject struct {
	Name string `bigquery:"column:Name"`
	Date string `bigquery:"column:Date"`
}

type RunTestSuit struct {
	Name string `bigquery:"column:Name"`
}

type RunChartVersion struct {
	Label     string    `bigquery:"column:Label"`
	Date      time.Time `bigquery:"column:Date"`
	Changeset string    `bigquery:"column:Changeset"`
	Branch    string    `bigquery:"column:Branch"`
}

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	db, err := gorm.Open("bigquery", "bigquery://go-bigquery-driver/playground")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&RunTestProject{})

	//unnestingExample(db)
	//unnestingExample2(db)
	hasTableExample(db)
	//miscExamples(db)

	defer db.Close()
	// Do Something with the DB

}

func unnestingExample(db *gorm.DB) {
	var versions []RunChartVersion
	err := db.Select("DISTINCT CONCAT(" +
		"CAST(sample.ProductVersion.MajorVersion AS STRING), '.'," +
		"CAST(sample.ProductVersion.MinorVersion AS STRING), '.'," +
		"CAST(sample.ProductVersion.RevisionVersion AS STRING)" +
		") as Label," +
		"sample.ProductVersion.Changeset," +
		"sample.EditorVersion.Branch," +
		"sample.ProductVersion.Date").Table("run_charts, UNNEST(Samples) as sample").Find(&versions).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, version := range versions {
		log.Printf("%s,%s\n", version.Label, version.Date)
	}
}

func unnestingExample2(db *gorm.DB) {
	var versions []RunChartVersion
	db.Select("DISTINCT '' as Label," +
		"sample.ProductVersion.Changeset," +
		"sample.ProductVersion.Branch," +
		"sample.ProductVersion.Date").Table("run_charts, UNNEST(Samples) as sample").Find(&versions)

	for _, version := range versions {
		log.Printf("%s,%s\n", version.Branch, version.Date)
	}
}

func miscExamples(db *gorm.DB) {
	var projects []RunTestProject
	var suits []RunTestSuit

	err := db.Find(&suits).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, suit := range suits {
		log.Println(suit.Name)
	}

	err = db.Not("Name", []string{"", "2D"}).Limit(2).Find(&projects).Error
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
}

func hasTableExample(db *gorm.DB) {
	var projects []RunTestProject
	if db.HasTable(projects) {
		log.Println("verified has table")
	} else {
		log.Println("could not verify table")
	}
}
