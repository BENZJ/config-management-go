package utils

import (
	"fmt"
	"time"
)

type Datetime time.Time

func (t Datetime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (dt Datetime) Format(layout string) string {
	return time.Time(dt).Format(layout)
}

func (dt Datetime) IsZero() bool {
	return time.Time(dt).IsZero()
}
