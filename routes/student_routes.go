package routes

import (
	"github.com/VictorFidelis/learning_gin/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterStudentRoutes(router *gin.Engine) {
	router.GET("/students", controllers.GetStudents)
	router.GET("/student/:id", controllers.GetStudentById)
	router.POST("/student", controllers.CreateStudant)
	router.DELETE("/student/:id", controllers.DeleteStudent)
	router.PUT("/student/:id", controllers.EditStudent)
	router.GET("/student/document/:document", controllers.GetStudentByDocument)
}
