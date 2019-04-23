package util

import "time"

func HandleTime(t string) (time.Time, error) {
	res, err := time.Parse("2006-01-02 15:04:05", t)
	return res, err
}
