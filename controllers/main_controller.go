package controllers

import (
	"net/http"

	"github.com/VictorFidelis/learning_gin/database"
	"github.com/VictorFidelis/learning_gin/models"
	"github.com/gin-gonic/gin"
)

func LoadHome(c *gin.Context) {
	var students []models.Student
	database.Db.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
