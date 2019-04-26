package routers

import (
	"farm/pkg/setting"
	v1 "farm/routers/api/v1"

	"github.com/gin-gonic/gin"

	_ "farm/docs"

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
	apiv1_calve := r.Group("/api/v1/calves")
	{
		apiv1_calve.GET("", v1.GetCalves)
		apiv1_calve.GET("/:cow_id", v1.GetCalve)
		apiv1_calve.POST("", v1.AddCalve)
		apiv1_calve.PUT("/:id", v1.UpdateCalve)
		apiv1_calve.DELETE("/:id", v1.DeleteCalve)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
