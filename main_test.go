package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/VictorFidelis/learning_gin/controllers"
	"github.com/VictorFidelis/learning_gin/database"
	"github.com/VictorFidelis/learning_gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var student models.Student

func CreateStudantTest() {
	student = models.Student{
		Name:     "Maria",
		Document: "1234568901",
	}
	database.Db.Create(&student)
}

func DeleteStudentTest() {
	database.Db.Delete(&student, student.ID)
}

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func TestGetStudents(t *testing.T) {
	database.ConnectDatabase()

	CreateStudantTest()
	defer DeleteStudentTest()

	routes := SetupTestRoutes()
	routes.GET("/students", controllers.GetStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	res := httptest.NewRecorder()
	routes.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "Deve obter o status 200")
}

func TestGetStudentsById(t *testing.T) {
	database.ConnectDatabase()

	CreateStudantTest()
	defer DeleteStudentTest()

	routes := SetupTestRoutes()
	routes.GET("/student/:id", controllers.GetStudentById)
	url := "/student/" + strconv.Itoa(int(student.ID))
	req, _ := http.NewRequest("GET", url, nil)
	res := httptest.NewRecorder()
	routes.ServeHTTP(res, req)

	var studentGet models.Student
	json.Unmarshal(res.Body.Bytes(), &studentGet)
	assert.Equal(t, student.Name, studentGet.Name, "Deve retornar o nome conforme a busca")
	assert.Equal(t, student.Document, studentGet.Document, "Deve retornar o documento conforme a busca")
	assert.Equal(t, http.StatusOK, res.Code, "Deve retornar status 200")
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectDatabase()

	CreateStudantTest()

	routes := SetupTestRoutes()
	routes.DELETE("/student/:id", controllers.DeleteStudent)
	url := "/student/" + strconv.Itoa(int(student.ID))
	req, _ := http.NewRequest("DELETE", url, nil)
	res := httptest.NewRecorder()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "Deve retornar status 200")
}

func TestEditStudent(t *testing.T) {
	database.ConnectDatabase()

	CreateStudantTest()
	defer DeleteStudentTest()

	routes := SetupTestRoutes()
	routes.PUT("/student/:id", controllers.EditStudent)

	studentPut := models.Student{
		Name:     student.Name,
		Document: "99999988845",
	}
	studentJson, _ := json.Marshal(studentPut)
	url := "/student/" + strconv.Itoa(int(student.ID))
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(studentJson))
	res := httptest.NewRecorder()
	routes.ServeHTTP(res, req)

	var studentEdited models.Student
	json.Unmarshal(res.Body.Bytes(), &studentEdited)

	assert.Equal(t, http.StatusOK, res.Code, "Deve retornar status 200")
	assert.Equal(t, studentPut.Name, studentEdited.Name, "Deve retornar o nome conforme o editado")
	assert.Equal(t, studentPut.Document, studentEdited.Document, "Deve retornar o documento conforme o editado")

}
