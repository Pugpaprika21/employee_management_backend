package employee

import (
	"github.com/Pugpaprika21/app/controller"
	"github.com/Pugpaprika21/app/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type employeeRouter struct {
	controller controller.IEmployeeController
}

func NewEmployeeRouter() *employeeRouter {
	return &employeeRouter{}
}

func (emp *employeeRouter) Register(e *echo.Echo, db *gorm.DB) {
	repository := repository.NewEmployeeRepository(db)
	emp.controller = controller.NewEmployeeController(repository)

	g := e.Group("/api/v1")
	g.POST("/createEmployee", emp.controller.CreateEmployee).Name = "createEmployee"
	g.GET("/getAllEmployee", emp.controller.GetAllEmployee).Name = "getAllEmployee"
	g.GET("/getEmployeeById/:id", emp.controller.GetEmployeeByID).Name = "getEmployeeById"
	g.PUT("/updateEmployeeById/:id", emp.controller.UpdateEmployeeByID).Name = "updateEmployeeById"
	g.DELETE("/deleteEmployeeById/:id", emp.controller.DeleteEmployeeByID).Name = "deleteEmployeeById"
}
