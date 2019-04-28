package util

import "regexp"

func CheckEmail(e string) bool {
	exp := regexp.MustCompile(`^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`)
	return exp.MatchString(e)
}
