package entity

import "database/sql/driver"

type MyBool bool

func (b MyBool) Value() (driver.Value, error) {
	result := make([]byte, 1)
	if b {
		result[0] = byte(1)
	} else {
		result[0] = 0
	}
	return result, nil
}
func (b MyBool) Scan(v interface{}) error {
	bytes := v.([]byte)

	if bytes[0] == 0 {
		b = false
	} else {
		b = true
	}
	return nil
}
