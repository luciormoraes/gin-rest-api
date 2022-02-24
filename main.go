package main

import (
	"github.com/luciormoraes/gin-rest-api/models"
	"github.com/luciormoraes/gin-rest-api/routes"
)

func main() {
	models.Students = []models.Student{
		{Name: "Lucio", IRD: "123123", DriverLicence: "987789"},
		{Name: "Suzana", IRD: "987789", DriverLicence: "123123"},
	}
	routes.HandleRequests()
}
