package dialect

import "github.com/jinzhu/gorm"

func init() {
	gorm.RegisterDialect(DialectName, &bigQueryDialect{})
}
