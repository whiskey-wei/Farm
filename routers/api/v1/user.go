package v1

import (
	"farm/model"
	"farm/pkg/e"
	"farm/pkg/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//log.Println(username, " ", password)
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	if username == "" || util.CheckUser(username) || password == "" || util.CheckUser(password) {
		code = e.ERROR_USER
	} else {
		if model.CheckUser(username, password) {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_USER_TOKEN
				log.Println(token)
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_USER
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
