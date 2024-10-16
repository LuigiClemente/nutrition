package utils

import (
	"strconv"
	"time"
)

func ParseAndValidateID(idStr string) (int, error) {
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func ParseTime(timeString string) (time.Time, error) {
	// Parse the time string using the provided layout
	layout := time.RFC3339
	parsedTime, err := time.Parse(layout, timeString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
