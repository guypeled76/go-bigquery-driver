package dialect

import "github.com/jinzhu/gorm"

func init() {
	gorm.RegisterDialect(Name, &bigQueryDialect{})
}
