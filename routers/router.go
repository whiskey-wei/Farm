package routers

import (
	"github.com/whiskey-wei/Farm/middleware/jwt"
	"github.com/whiskey-wei/Farm/pkg/setting"
	v1 "github.com/whiskey-wei/Farm/routers/api/v1"

	"github.com/gin-gonic/gin"

	_ "github.com/whiskey-wei/Farm/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath /api/v1/

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	//获取Token
	r.GET("/user", v1.GetUserToken)
	//注册，添加用户
	r.POST("/user", v1.AddUser)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//calverecord
		apiv1.GET("/calves", v1.GetCalves)
		apiv1.GET("/calves/:cow_id", v1.GetCalve)
		apiv1.POST("/calves", v1.AddCalve)
		apiv1.PUT("/calves/:id", v1.UpdateCalve)
		apiv1.DELETE("/calves/:id", v1.DeleteCalve)

		apiv1.PUT("/user", v1.UpdateUser)
		apiv1.GET("/user", v1.GetUserSelf)
	}

	apiAdmin := r.Group("/admin")
	apiAdmin.Use(jwt.JwtAdmin())
	{
		apiAdmin.GET("/user", v1.GetUser)
		apiAdmin.GET("/users", v1.GetUsers)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
