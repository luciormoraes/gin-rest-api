package models

type Student struct {
	// ID   int    `json:"ID"`
	Name          string `json:"name"`
	IRD           string `json:"ird"`
	DriverLicence string `json:"driverLicence"`
}

var Students []Student
