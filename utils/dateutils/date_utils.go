package dateutils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05"
)

//GetNow returns current time in UTC
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString gets date in predefined format
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
