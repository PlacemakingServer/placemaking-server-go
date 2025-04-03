package utils

import (
	"time"
)

func ParseTimestamp(t string) (string, error) {
	parsedTime, err := time.Parse("2006-01-02T15:04:05.999999", t)
	if err != nil {
		return "", err
	}
	return parsedTime.Format(time.RFC3339), nil
}
