package util

import (
	"github.com/whiskey-wei/Farm/pkg/setting"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	//fmt.Println("page:", page, " res:", result)
	return result
}
