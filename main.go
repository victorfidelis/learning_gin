package main

import (
	"github.com/VictorFidelis/learning_gin/database"
	"github.com/VictorFidelis/learning_gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDatabase()

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	routes.RegisterMainRoutes(router)
	routes.RegisterStudentRoutes(router)

	router.Run()
}
