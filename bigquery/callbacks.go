package bigquery

import (
	"github.com/guypeled76/go-bigquery-driver/adaptor"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
)

func initializeCallbacks(db *gorm.DB) {

	// register callbacks
	callbacks.RegisterDefaultCallbacks(db, &callbacks.Config{
		WithReturning: true,
	})

	c := &bigQueryCallbacks{db}

	queryCallback := db.Callback().Query()
	queryCallback.Replace("gorm:query", c.queryCallback)
}

type bigQueryCallbacks struct {
	root *gorm.DB
}

func (c *bigQueryCallbacks) queryCallback(db *gorm.DB) {
	if !db.DryRun {

		db.Statement.Context = adaptor.SetSchemaAdaptor(db.Statement.Context, &bigQuerySchemaAdaptor{
			db.Statement.Schema,
			c.root,
		})
	}

	callbacks.Query(db)
}
