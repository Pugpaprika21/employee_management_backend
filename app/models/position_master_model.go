package models

import "gorm.io/gorm"

type PositionMaster struct {
	gorm.Model
	Name             string  `gorm:"not null"`
	Description      string  `gorm:"not null"`
	Salary           float64 `gorm:"not null"`
	Status           string
	Location         string
	Responsibilities string // ความรับผิดชอบ
}
