package utils

import "time"

func ParseDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02" // Format YYYY-MM-DD
	return time.Parse(layout, dateStr)
}
