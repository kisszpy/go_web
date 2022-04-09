package common

import (
	"crypto/md5"
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

type MyTime time.Time

func (this *MyTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*this = MyTime(now)
	return
}

func (this MyTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(this)
	if tt.IsZero() {
		emptyStr, _ := json.Marshal("")
		return emptyStr, nil
	}
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (this MyTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tt := time.Time(this)
	if tt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tt, nil
}

func (this *MyTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*this = MyTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (this MyTime) IsZero() bool {
	tt := time.Time(this)
	if tt.IsZero() {
		return true
	}
	return false
}

func (this MyTime) String() string {
	tt := time.Time(this)
	return tt.String()
}

func (this MyTime) Format() string {
	tt := time.Time(this)
	return tt.Format(timeFormat)
}

func (this MyTime) After(t MyTime) bool {
	tt := time.Time(this)
	return tt.After(time.Time(t))
}

func (this MyTime) Before(t MyTime) bool {
	tt := time.Time(this)
	return tt.Before(time.Time(t))
}

func (this MyTime) Parse(t string) (MyTime, error) {
	tt, err := time.Parse(timeFormat, t)
	return MyTime(tt), err
}

func Md5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
