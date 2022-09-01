package utils

import "time"

// StringToDate converts string to date
func StringToDate(date string, layout string) (time.Time, error) {
	return time.Parse(layout, date)
}
