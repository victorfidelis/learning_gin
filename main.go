package main

import (
	"github.com/VictorFidelis/learning_gin/database"
	"github.com/VictorFidelis/learning_gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDatabase()

	routes.RegisterStudentRoutes(router)
	router.Run()
}
