package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/whiskey-wei/Farm/model"
	"github.com/whiskey-wei/Farm/pkg/e"
	"github.com/whiskey-wei/Farm/pkg/setting"
	"github.com/whiskey-wei/Farm/pkg/util"

	"github.com/Unknwon/com"

	"github.com/gin-gonic/gin"
)

//登录获取token,username,password在路径中
func GetUserToken(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//role := com.StrTo(c.Query("role")).MustInt()
	//log.Println(username, " ", password)
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})

	role := model.CheckUser(username, password)
	if role != 0 {
		token, err := util.GenerateToken(username, password, role)
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

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//注册，添加用户，表单提交，所有的数据都是必填项，只能注册普通用户role为1
func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	role := 1
	telephonenumber := c.PostForm("telephone_number")
	email := c.PostForm("e_mail")
	fmt.Println(username + " " + password)
	code := e.INVALID_PARAMS

	if model.ExistUserByUsername(username) {
		code = e.ERROR_EXIST
	} else {
		model.AddUser(model.User{
			Username:        username,
			Password:        password,
			Role:            role,
			TelephoneNumber: telephonenumber,
			EMail:           email,
		})
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

//用户获取自身信息
func GetUserSelf(c *gin.Context) {
	token := c.Query("token")
	claim, _ := util.ParseToken(token)
	username := claim.Username

	code := e.INVALID_PARAMS
	data := make(map[string]interface{})

	user, ok := model.GetUserByUsername(username)
	if ok {
		code = e.SUCCESS
		data["user"] = user
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//更新数据
func UpdateUser(c *gin.Context) {
	token := c.Query("token")
	claim, _ := util.ParseToken(token)
	username := claim.Username
	password := c.PostForm("password")
	role := com.StrTo(c.PostForm("role")).MustInt()
	telephonenumber := c.PostForm("telephone_number")
	email := c.PostForm("e_mail")

	code := e.INVALID_PARAMS

	if (username != "" && !util.CheckUser(username)) ||
		password != "" && !util.CheckUser(password) ||
		telephonenumber != "" && !util.CheckTelepthoneNumber(telephonenumber) ||
		email != "" && !util.CheckEmail(email) ||
		role != 0 && role != 1 && role != 2 {
		code = e.INVALID_PARAMS
		log.Println("用户名密码格式出错")
	} else {
		if model.UpdateUser(model.User{
			Password:        password,
			Role:            role,
			TelephoneNumber: telephonenumber,
			EMail:           email,
		}) {
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

//通过username查询信息,管理员操作
func GetUser(c *gin.Context) {
	token := c.Query("token")
	username := c.Query("username")
	claim, _ := util.ParseToken(token)
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	if claim.Role != 2 {
		code = e.ERROR_PERMISSION_DENY
	} else {
		if util.CheckUser(username) {
			user, ok := model.GetUserByUsername(username)
			if ok {
				data["data"] = user
				code = e.SUCCESS
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//获取所有用户信息，管理员操作
func GetUsers(c *gin.Context) {
	token := c.Query("token")
	claim, _ := util.ParseToken(token)
	code := e.INVALID_PARAMS
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	total := 0
	if claim.Role != 2 {
		code = e.ERROR_PERMISSION_DENY
	} else {
		total = model.GetUserTotal(maps)
		data["total"] = total
		if total > 0 {
			users, ok := model.GetUsers(util.GetPage(c), setting.PageSize, maps)
			if ok {
				data["lists"] = users
				code = e.SUCCESS
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  data,
		"total": total,
	})
}

//管理员修改用户信息
