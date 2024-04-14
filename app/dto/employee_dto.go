package dto

import "time"

type EmployeeRequestBody struct {
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	PositionID   uint      `json:"position_id"`
	DepartmentID uint      `json:"department_id"`
	StartDate    time.Time `json:"start_date"`
	Salary       float64   `json:"salary"`
	Address      string    `json:"address"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Gender       string    `json:"gender"`
	DateOfBirth  time.Time `json:"date_of_birth"`
}

type EmployeeFetchRecord struct {
	ID           uint    `gorm:"id"`
	Name         string  `gorm:"name"`
	Surname      string  `gorm:"surname"`
	PositionID   uint    `gorm:"position_id"`
	DepartmentID uint    `gorm:"department_id"`
	Salary       float64 `gorm:"salary"`
	Address      string  `gorm:"address"`
	PhoneNumber  string  `gorm:"phone_number"`
	Email        string  `gorm:"email"`
	Gender       string  `gorm:"gender"`
	CreatedAt    string  `gorm:"created_at"`
	UpdatedAt    string  `gorm:"updated_at"`
	DeletedAt    string  `gorm:"deleted_at"`
}

type EmployeeResponesData struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	Surname        string  `json:"surname"`
	PositionID     uint    `json:"position_id"`
	DepartmentID   uint    `json:"department_id"`
	PositionName   string  `json:"position_name"`
	DepartmentName string  `json:"department_name"`
	Salary         float64 `json:"salary"`
	Address        string  `json:"address"`
	PhoneNumber    string  `json:"phone_number"`
	Email          string  `json:"email"`
	Gender         string  `json:"gender"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	DeletedAt      string  `json:"deleted_at"`
}

type EmployeeRespones struct {
	Message    string      `json:"message"`
	StatusBool bool        `json:"status_bool"`
	Data       interface{} `json:"data,omitempty"`
}
