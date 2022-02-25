package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luciormoraes/gin-rest-api/database"
	"github.com/luciormoraes/gin-rest-api/models"
)

func AllStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Salutation(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
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
