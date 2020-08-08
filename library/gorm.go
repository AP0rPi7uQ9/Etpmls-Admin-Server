package library

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// https://www.jianshu.com/p/e30f299f7b6c
// Date formatted to seconds
// 日期格式化到秒
type TimeSecond struct {
	time.Time
}

func (t TimeSecond) MarshalJSON() ([]byte, error) {
	// tune := fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

// Value insert timestamp into mysql need this function.
func (t TimeSecond) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *TimeSecond) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeSecond{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}



// https://www.jianshu.com/p/e30f299f7b6c
// Date formatted to days
// 日期格式化到天
type TimeDay struct {
	time.Time
}

func (t TimeDay) MarshalJSON() ([]byte, error) {
	// tune := fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))
	tune := t.Format(`"2006-01-02"`)
	return []byte(tune), nil
}

// Value insert timestamp into mysql need this function.
func (t TimeDay) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *TimeDay) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeDay{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}