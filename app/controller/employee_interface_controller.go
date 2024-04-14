package controller

import "github.com/labstack/echo/v4"

type IEmployeeController interface {
	CreateEmployee(c echo.Context) error
	GetAllEmployee(c echo.Context) error
	GetEmployeeByID(c echo.Context) error
	UpdateEmployeeByID(c echo.Context) error
	DeleteEmployeeByID(c echo.Context) error
}
