package bigquery

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/migrator"
	"gorm.io/gorm/schema"
)

type Migrator struct {
	migrator.Migrator
}

func (m Migrator) CurrentDatabase() (name string) {
	m.DB.Raw("SELECT CURRENT_DATABASE()").Row().Scan(&name)
	return
}

func (m Migrator) BuildIndexOptions(opts []schema.IndexOption, stmt *gorm.Statement) (results []interface{}) {
	return
}

func (m Migrator) HasIndex(value interface{}, name string) bool {
	return false
}

func (m Migrator) CreateIndex(value interface{}, name string) error {
	return errors.New("CreateIndex is unsupported")
}

func (m Migrator) RenameIndex(value interface{}, oldName, newName string) error {
	return errors.New("RenameIndex is unsupported")
}

func (m Migrator) DropIndex(value interface{}, name string) error {
	return errors.New("DropIndex is unsupported")
}

func (m Migrator) HasTable(value interface{}) bool {
	var count int64
	m.RunWithValue(value, func(stmt *gorm.Statement) error {
		return m.DB.Raw("SELECT count(*) FROM information_schema.tables WHERE table_name = ?", stmt.Table).Row().Scan(&count)
	})

	return count > 0
}

func (m Migrator) DropTable(values ...interface{}) error {
	values = m.ReorderModels(values, false)
	tx := m.DB.Session(&gorm.Session{})
	for i := len(values) - 1; i >= 0; i-- {
		if err := m.RunWithValue(values[i], func(stmt *gorm.Statement) error {
			return tx.Exec("DROP TABLE IF EXISTS ?", clause.Table{Name: stmt.Table}).Error
		}); err != nil {
			return err
		}
	}
	return nil
}

func (m Migrator) HasColumn(value interface{}, field string) bool {
	var count int64
	m.RunWithValue(value, func(stmt *gorm.Statement) error {
		name := field
		if field := stmt.Schema.LookUpField(field); field != nil {
			name = field.DBName
		}

		return m.DB.Raw(
			"SELECT count(*) FROM INFORMATION_SCHEMA.columns WHERE table_schema = CURRENT_SCHEMA() AND table_name = ? AND column_name = ?",
			stmt.Table, name,
		).Row().Scan(&count)
	})

	return count > 0
}

func (m Migrator) HasConstraint(value interface{}, name string) bool {
	var count int64
	m.RunWithValue(value, func(stmt *gorm.Statement) error {
		return m.DB.Raw(
			"SELECT count(*) FROM INFORMATION_SCHEMA.table_constraints WHERE table_schema = CURRENT_SCHEMA() AND table_name = ? AND constraint_name = ?",
			stmt.Table, name,
		).Row().Scan(&count)
	})

	return count > 0
}
