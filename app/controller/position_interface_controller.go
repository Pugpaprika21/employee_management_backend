package controller

import "github.com/labstack/echo/v4"

type IPositionController interface {
	CreatePosition(c echo.Context) error
	GetAllPosition(c echo.Context) error
	GetPositionByID(c echo.Context) error
	UpdatePositionByID(c echo.Context) error
	DeletePositionByID(c echo.Context) error
}
