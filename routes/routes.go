package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luciormoraes/gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.AllStudents)
	r.GET("/:name", controllers.Salutation)
	r.POST("/student", controllers.CreateStudent)
	r.GET("/students/:id", controllers.SearchStudentByID)
	r.Run()
}
