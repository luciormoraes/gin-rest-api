package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luciormoraes/gin-rest-api/database"
	"github.com/luciormoraes/gin-rest-api/models"
)

func AllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func Salutation(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"API says": "Hey " + name + ", all good?",
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, &student)
}

func SearchStudentByID(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Student does not exist in database",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}
