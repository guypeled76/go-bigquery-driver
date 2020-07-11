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

	rootDB := db

	queryCallback := db.Callback().Query()
	queryCallback.Replace("gorm:query", func(db *gorm.DB) {
		if !db.DryRun {

			db.Statement.Context = adaptor.SetSchemaAdaptor(db.Statement.Context, &bigQuerySchemaAdaptor{
				db.Statement.Schema,
				rootDB,
			})
		}

		callbacks.Query(db)
	})
}
