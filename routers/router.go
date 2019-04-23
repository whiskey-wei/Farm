package routers

import (
	"farm/pkg/setting"
	v1 "farm/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/calves", v1.GetCalves)
		apiv1.GET("/calves/:id", v1.GetCalve)
		apiv1.POST("/calves", v1.AddCalve)
		apiv1.PUT("/calves/:id", v1.UpdateCalve)
		apiv1.DELETE("/calves/:id", v1.DeleteCalve)
	}
	return r
}
