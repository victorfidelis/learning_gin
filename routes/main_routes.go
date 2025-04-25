package routes

import (
	"github.com/VictorFidelis/learning_gin/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterMainRoutes(router *gin.Engine) {
	router.GET("/", controllers.LoadHome)
	router.NoRoute(controllers.RouteNotFound)
}
