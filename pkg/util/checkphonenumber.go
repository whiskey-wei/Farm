package util

import "regexp"

//检查中国地区电话号码
func CheckTelepthoneNumber(number string) bool {
	exp := regexp.MustCompile(`^1(3|4|5|7|8)\d{9}$`)
	return exp.MatchString(number)
}
