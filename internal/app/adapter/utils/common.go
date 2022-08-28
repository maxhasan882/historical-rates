package utils

import "time"

func StringToDate(date string, layout string) (time.Time, error) {
	t, err := time.Parse(layout, date)
	return t, err
}
