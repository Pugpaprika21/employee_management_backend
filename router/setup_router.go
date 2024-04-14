package router

import (
	"github.com/Pugpaprika21/router/department"
	"github.com/Pugpaprika21/router/employee"
	"github.com/Pugpaprika21/router/position"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type serverRouter struct {
	router IRouter
}

func NewServerRouter() *serverRouter {
	return &serverRouter{}
}

func (s *serverRouter) Setup(e *echo.Echo, db *gorm.DB) {
	s.router = employee.NewEmployeeRouter()
	s.router.Register(e, db)

	s.router = department.NewDepartmentRouter()
	s.router.Register(e, db)

	s.router = position.NewPositionRouter()
	s.router.Register(e, db)
}
