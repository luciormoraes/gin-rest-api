package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name          string `json:"name" validate:"nonzero"`
	IRD           string `json:"ird" validate:"len=6,regexp=^[0-9]*$"`
	DriverLicence string `json:"driverLicence" validate:"len=7,regexp=^[0-9]*$"`
}

func ValidateStudent(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
