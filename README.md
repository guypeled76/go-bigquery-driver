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

For gorm

```go
package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/guypeled76/go-bigquery-driver/gorm/dialect"
    "log"
)

func main() {
    db, err := gorm.Open("bigquery", 
        "bigquery://projectid/dataset")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() 
    // Do Something with the DB

}
```

