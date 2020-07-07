package scanner

func Scan(source interface{}, dest interface{}) error {
	return db.Where("SOURCE = ?", source).Scan(dest).Error
}
