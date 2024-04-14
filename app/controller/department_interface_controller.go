package controller

import "github.com/labstack/echo/v4"

type IDepartmentController interface {
	CreateDepartment(c echo.Context) error
	GetAllDepartment(c echo.Context) error
	GetDepartmentByID(c echo.Context) error
	UpdateDepartmentByID(c echo.Context) error
	DeleteDepartmentByID(c echo.Context) error
}
