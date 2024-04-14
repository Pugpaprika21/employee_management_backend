package models

import (
	"gorm.io/gorm"
)

type DepartmentMaster struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Location string `gorm:"not null"`
}
