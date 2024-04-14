package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type IRouter interface {
	Register(e *echo.Echo, db *gorm.DB)
}
