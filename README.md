#  BigQuery SQL Driver & GORM Dialect for Golang
This is an implementation of the BigQuery Client as a database/sql/driver for easy integration and usage.


# Goals of project

This module implements a BigQuery SQL driver and GORM dialect. 

# Usage

As this is using the Google Cloud Go SDK, you will need to have your credentials available
via the GOOGLE_APPLICATION_CREDENTIALS environment variable point to your credential JSON file.

## Vanilla *sql.DB usage

Just like any other database/sql driver you'll need to import it 

```go
package main

import (
    "database/sql"
    _ "github.com/guypeled76/go-bigquery-driver/driver"
    "log"
)

func main() {
    db, err := sql.Open("bigquery", 
        "bigquery://projectid/dataset")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() 
    // Do Something with the DB

}
```

## Gorm Usage

Opening a Gorm bigquery db

```go
package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/guypeled76/go-bigquery-driver/gorm/dialect"
    "log"
)

func main() {
    // You can also use the location format: "bigquery://projectid/location/dataset"
    db, err := gorm.Open("bigquery", "bigquery://projectid/dataset")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() 
    // Do Something with the DB

}
```


Using gorm with a BigQuery query that has a struct

```go
package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/guypeled76/go-bigquery-driver/gorm/dialect"
    "github.com/guypeled76/go-bigquery-driver/gorm/scanner"
    "log"
)


type ComplexRecord struct {
	Name   string           `gorm:"column:Name"`
	Record ComplexSubRecord `gorm:"column:Record"`
}

type ComplexSubRecord struct {
	Name string `gorm:"column:Name"`
	Age  int    `gorm:"column:Age"`
}

// Scan overrides the default GORM behavior for structs and allows to 
// get results from a big query query which has structs in it.
func (record *ComplexSubRecord) Scan(value interface{}) error {
	return scanner.Scan(value, record)
}


func main() {
    db, err := gorm.Open("bigquery",  "bigquery://projectid/dataset")
    if err != nil {
        log.Fatal(err)
    }

    var records []ComplexRecord
    
    // Delete complex record table if exists
    db.DropTableIfExists(&ComplexRecord{})

    // Make sure we have a complex_records table
    db.AutoMigrate(&ComplexRecord{})
    
    // Insert new records to table
    db.Create(&ComplexRecord{Name: "test", Record: ComplexSubRecord{Name: "dd", Age: 1}})
    db.Create(&ComplexRecord{Name: "test2", Record: ComplexSubRecord{Name: "dd2", Age: 444}})

    // Select records from table
    db.Order("Name").Find(&records)


    defer db.Close() 
    // Do Something with the DB

}
```

Using gorm with a BigQuery query that has an array

```go
package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/guypeled76/go-bigquery-driver/gorm/dialect"
    "github.com/guypeled76/go-bigquery-driver/gorm/scanner"
    "log"
)

type ArrayRecord struct {
	Name    string            `gorm:"column:Name"`
	Records ArrayRecordRecord `gorm:"column:Records"`
}

type ArrayRecordRecord []ComplexSubRecord

func (record *ArrayRecordRecord) Scan(value interface{}) error {
	return scanner.Scan(value, record)
}

type ComplexSubRecord struct {
	Name string `gorm:"column:Name"`
	Age  int    `gorm:"column:Age"`
}

func (record *ComplexSubRecord) Scan(value interface{}) error {
	return scanner.Scan(value, record)
}



func main() {
    db, err := gorm.Open("bigquery",  "bigquery://projectid/dataset")
    if err != nil {
        log.Fatal(err)
    }

    var records []ArrayRecord
    
    // Delete array_records table if exists
    db.DropTableIfExists(&ArrayRecord{})

    // Make sure we have an array_records table
    db.AutoMigrate(&ArrayRecord{})
    
    // Insert new records to table
    db.Create(&ArrayRecord{Name: "test", Records: ArrayRecordRecord{{Name: "dd", Age: 1}, {Name: "dd1", Age: 1}}})
    db.Create(&ArrayRecord{Name: "test2", Records: ArrayRecordRecord{{Name: "dd2", Age: 444}, {Name: "dd3", Age: 1}}})

    // Select records from table ordered by name
    db.Order("Name").Find(&records)


    defer db.Close() 
    // Do Something with the DB

}
```

Using gorm with a BigQuery query that uses unnest

```go
package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/guypeled76/go-bigquery-driver/gorm/dialect"
    "github.com/guypeled76/go-bigquery-driver/gorm/scanner"
    "log"
)

type Version struct {
	Label string `gorm:"column:Label"`
}

func main() {
    db, err := gorm.Open("bigquery",  "bigquery://projectid/dataset")
    if err != nil {
        log.Fatal(err)
    }

    var versions []Version
    
    query := db.Table("charts, UNNEST(Samples) as sample")

    query = query.Select("DISTINCT CONCAT(" +
        "CAST(sample.MajorVersion AS STRING), '.'," +
        "CAST(sample.MinorVersion AS STRING), '.'," +
        "CAST(sample.RevisionVersion AS STRING)" +
        ") as Label")
    
    err := query.Find(&versions).Error
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close() 
    // Do Something with the DB

}
```