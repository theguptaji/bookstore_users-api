package date_utils

import "time"

const (
	apiDateLayout = "2016-01-02T15:04:05Z"
	apiDbLayout   = "2016-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDbFormat() string {
	return GetNow().Format(apiDbLayout)
}
