package routes

import (
	"github.com/VictorFidelis/learning_gin/controllers"
	docs "github.com/VictorFidelis/learning_gin/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterMainRoutes(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/", controllers.LoadHome)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.NoRoute(controllers.RouteNotFound)
}
