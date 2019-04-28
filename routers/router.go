package routers

import (
	"farm/middleware/jwt"
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

	//获取Token
	r.GET("/user", v1.GetUser)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//calverecord
		apiv1.GET("/calves", v1.GetCalves)
		apiv1.GET("/calves/:cow_id", v1.GetCalve)
		apiv1.POST("/calves", v1.AddCalve)
		apiv1.PUT("/calves/:id", v1.UpdateCalve)
		apiv1.DELETE("/calves/:id", v1.DeleteCalve)
	}
	/*apiv1_breed := r.Group("/api/v1/breeds")
	{
		apiv1_breed.GET("", v1.GetBreeds)
		apiv1_breed.GET("/:cow_id", v1.GetBreed)
		apiv1_breed.POST("", v1.AddBreed)
		apiv1_breed.PUT("/:id", v1.UpdateBreed)
		apiv1_breed.DELETE("/:id", v1.DeleteBreed)
	}*/
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
