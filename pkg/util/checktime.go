package util

import (
	"regexp"
	"time"
)

func HandleTime(t string) (time.Time, error) {
	res, err := time.Parse("2006-01-02 15:04:05", t)
	return res, err
}

func HandleUnixTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

//正则判断时间格式
func CheckTime(t string) bool {
	exp := regexp.MustCompile(`[12]\d{3}-[01]\d-[0123]\d [012]\d:[012345]\d:[012345]\d`)
	return exp.MatchString(t)
}
