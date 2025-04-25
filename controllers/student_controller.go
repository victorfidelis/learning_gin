package controllers

import (
	"net/http"
	"strconv"

	"github.com/VictorFidelis/learning_gin/database"
	"github.com/VictorFidelis/learning_gin/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	database.Db.Find(&students)
	c.JSON(http.StatusOK, students)
}

// GetStudentById godoc
// @Summary      Apresentar um aluno
// @Description  Busca um aluno através do seu id
// @Tags         student
// @Produce      json
// @Param        id   path      int  true  "Id do aluno"
// @Success      200  {object}  models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /student/{id} [get]
func GetStudentById(c *gin.Context) {
	textId := c.Params.ByName("id")
	id, err := strconv.Atoi(textId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var student models.Student
	database.Db.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

// CreateStudant godoc
// @Summary      Criar um novo aluno
// @Description  Cria um novo aluno através do json passado no body
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        id   body      models.Student  true  "Modelo do aluno"
// @Success      200  {object}  models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /student [post]
func CreateStudant(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := student.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.Db.Create(&student)
	c.JSON(http.StatusCreated, student)

}

func DeleteStudent(c *gin.Context) {
	textId := c.Params.ByName("id")
	id, err := strconv.Atoi(textId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var student models.Student
	database.Db.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}

	database.Db.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Aluno removido",
	})
}

func EditStudent(c *gin.Context) {
	textId := c.Params.ByName("id")
	id, err := strconv.Atoi(textId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var student models.Student

	database.Db.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := student.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.Db.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func GetStudentByDocument(c *gin.Context) {
	document := c.Params.ByName("document")

	var student models.Student
	database.Db.Where(&models.Student{Document: document}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}
