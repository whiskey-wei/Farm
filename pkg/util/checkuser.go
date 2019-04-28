package util

import "regexp"

func CheckUser(str string) bool {
	exp := regexp.MustCompile(`/^[a-zA-Z0-9_-]{4,16}$/`)
	return exp.MatchString(str)
}
