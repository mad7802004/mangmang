package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type JSONTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*t = JSONTime{Time: local}
	return err
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}



type JSONDate struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d
func (t JSONDate) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *JSONDate) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(`"2006-01-02"`, string(data), time.Local)
	*t = JSONDate{Time: local}
	return err
}

// Scan valueof time.Time
func (t *JSONDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
