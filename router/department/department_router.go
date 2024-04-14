package department

import (
	"github.com/Pugpaprika21/app/controller"
	"github.com/Pugpaprika21/app/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type departmentRouter struct {
	controller controller.IDepartmentController
}

func NewDepartmentRouter() *departmentRouter {
	return &departmentRouter{}
}

func (d *departmentRouter) Register(e *echo.Echo, db *gorm.DB) {
	repository := repository.NewDepartmentRepository(db)
	d.controller = controller.NewDepartmentController(repository)

	g := e.Group("/api/v1")
	g.POST("/createDepartment", d.controller.CreateDepartment).Name = "createDepartment"
	g.GET("/getAllDepartment", d.controller.GetAllDepartment).Name = "getAllDepartment"
	g.GET("/getDepartmentById/:id", d.controller.GetDepartmentByID).Name = "getDepartmentById"
	g.PUT("/updateDepartmentById/:id", d.controller.UpdateDepartmentByID).Name = "updateDepartmentById"
	g.DELETE("/deleteDepartmentById/:id", d.controller.DeleteDepartmentByID).Name = "deleteDepartmentById"
}
