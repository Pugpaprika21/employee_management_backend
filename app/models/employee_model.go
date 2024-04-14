package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name         string `gorm:"type:varchar(200);not null"`
	Surname      string `gorm:"type:varchar(200);not null"`
	PositionID   uint
	DepartmentID uint
	Salary       float64          `gorm:"not null"`
	Address      string           `gorm:"not null"`
	PhoneNumber  string           `gorm:"not null"`
	Email        string           `gorm:"not null"`
	Gender       string           `gorm:"not null"`
	Position     PositionMaster   `gorm:"foreignKey:PositionID"`
	Department   DepartmentMaster `gorm:"foreignKey:DepartmentID"`
}
