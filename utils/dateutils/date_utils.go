package dateutils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05"
	apiDbLayout   = "2006-01-02 15:04:05"
)

//GetNow returns current time in UTC
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString gets date in predefined format
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

//GetNowDBFormat gets date for db persistence
func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
