package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name          string `json:"name"`
	IRD           string `json:"ird"`
	DriverLicence string `json:"driverLicence"`
}

var Students []Student
