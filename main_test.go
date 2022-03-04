package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/luciormoraes/gin-rest-api/controllers"
	"github.com/luciormoraes/gin-rest-api/database"
	"github.com/luciormoraes/gin-rest-api/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func RouteSetupTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "test", IRD: "123456", DriverLicence: "1234567"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

// func TestFailure(t *testing.T) {
// 	t.Fatalf("Test has failed")
// }

func TestStatusCodeOK(t *testing.T) {
	r := RouteSetupTest()
	r.GET("/:name", controllers.Salutation)
	req, _ := http.NewRequest("GET", "/gin", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	// if response.Code != http.StatusOK {
	// 	t.Fatalf("Status error: Received value is %d. Expected was %d", response.Code, http.StatusOK)
	// }
	assert.Equal(t, http.StatusOK, response.Code, "Status error: Received value is not 200")
	mockResponse := `{"API says":"Hey gin, all good?"}`
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, mockResponse, string(responseBody), "Must be equal")
}

func TestGetAllStudentsHandler(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RouteSetupTest()
	r.GET("/students", controllers.AllStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Status error: Received value is not 200")
}

func TestSearchStudentByIRDHandler(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RouteSetupTest()
	r.GET("/students/ird/:ird", controllers.SearchStudentByIRD)
	req, _ := http.NewRequest("GET", "/students/ird/123456", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Status error: Received value is not 200")
}

func TestSearchStudentByID(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RouteSetupTest()
	r.GET("/students/:id", controllers.SearchStudentByID)

	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)
	assert.Equal(t, "test", studentMock.Name)
	assert.Equal(t, "123456", studentMock.IRD)
	assert.Equal(t, "1234567", studentMock.DriverLicence)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteStudentByIDHandler(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()

	r := RouteSetupTest()
	r.DELETE("/students/:id", controllers.DeleteStudentByID)
	path := "/students/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditStudentHandler(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := RouteSetupTest()
	r.PATCH("/students/:id", controllers.EditStudent)

	student := models.Student{Name: "test", IRD: "101010", DriverLicence: "1010107"}
	studentJson, _ := json.Marshal(student)

	path := "/students/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(studentJson))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)
	assert.Equal(t, "test", studentMock.Name)
	assert.Equal(t, "101010", studentMock.IRD)
	assert.Equal(t, "1010107", studentMock.DriverLicence)
	// assert.Equal(t, http.StatusOK, response.Code)
}
