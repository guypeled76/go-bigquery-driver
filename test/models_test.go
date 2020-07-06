package test

type SimpleTestRecord struct {
	Name string `gorm:"column:Name"`
}

type ComplexRecord struct {
	Name   string           `gorm:"column:Name"`
	Record ComplexSubRecord `gorm:"column:Record"`
}

type ComplexSubRecord struct {
	Name string `gorm:"column:Name"`
	Age  int    `gorm:"column:Age"`
}

func (record *ComplexSubRecord) Scan(value interface{}) error {
	return nil
}
