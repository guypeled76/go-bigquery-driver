package utils

import "database/sql/driver"

func GetValue(value driver.Value) interface{} {
	named, ok := value.(driver.NamedValue)
	if ok {
		return named.Value
	}
	return value
}

func GetValueAt(values []driver.Value, index int) interface{} {
	if len(values) <= index {
		return nil
	}
	return GetValue(values[index])
}

func GetStringValueAt(values []driver.Value, index int) string {
	return GetValueAt(values, index).(string)
}
