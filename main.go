package main

import (
	"github.com/luciormoraes/gin-rest-api/database"
	"github.com/luciormoraes/gin-rest-api/routes"
)

func main() {
	database.ConnectDB()

	routes.HandleRequests()
}
