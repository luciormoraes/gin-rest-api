package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luciormoraes/gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/students", controllers.AllStudents)
	r.GET("/:name", controllers.Salutation)
	r.POST("/student", controllers.CreateStudent)
	r.GET("/students/:id", controllers.SearchStudentByID)
	r.DELETE("/students/:id", controllers.DeleteStudentByID)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/ird/:ird", controllers.SearchStudentByIRD)
	r.GET("/index", controllers.ShowIndexPage)
	r.NoRoute(controllers.RouteNotFound)
	r.Run()
}
