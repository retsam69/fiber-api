package model

// import (
// 	"fmt"
// 	"strings"
// 	"time"
// )

// const (
// 	DATE_LAYOUT      = "2006-01-02"
// 	DATETIME_LAYOUT  = "2006-01-02T15:04:05"
// 	TIMESTAMP_LAYOUT = "2006-01-02T15:04:05.000+07:00"
// )

// type Date time.Time

// func (m *Date) UnmarshalJSON(b []byte) error {
// 	s := strings.Trim(string(b), `"`)
// 	tt, err := time.ParseInLocation(DATE_LAYOUT, s, time.Local)
// 	m.FromTime(tt)
// 	return err
// }

// func (ct Date) MarshalJSON() ([]byte, error) {
// 	return []byte(ct.String()), nil
// }

// func (ct *Date) String() string {
// 	t := time.Time(*ct)
// 	return fmt.Sprintf("%q", t.Format(DATE_LAYOUT))
// }

// func (t Date) GetTime() time.Time {
// 	return time.Time(t)
// }

// func (t *Date) FromTime(ti time.Time) {
// 	*t = Date(ti)
// }

// func (m *Date) UnmarshalCSV(b []byte) error {
// 	s := strings.Trim(string(b), `"`)
// 	tt, err := time.ParseInLocation(DATE_LAYOUT, s, time.Local)
// 	m.FromTime(tt)
// 	return err
// }

// func (ct Date) MarshalCSV() ([]byte, error) {
// 	t := time.Time(ct)
// 	return []byte(t.Format(DATE_LAYOUT)), nil
// }
