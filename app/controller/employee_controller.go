package controller

import (
	"net/http"

	"github.com/Pugpaprika21/app/dto"
	"github.com/Pugpaprika21/app/repository"
	"github.com/labstack/echo/v4"
)

type employeeController struct {
	repository repository.IEmployeeRepository
}

func NewEmployeeController(repository *repository.EmployeeRepository) *employeeController {
	return &employeeController{
		repository: repository,
	}
}

func (emp *employeeController) CreateEmployee(c echo.Context) error {
	var body *dto.EmployeeRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, dto.EmployeeRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	err := emp.repository.Save(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.EmployeeRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.EmployeeRespones{
		Message:    "save employee success",
		StatusBool: true,
	})
}

func (emp *employeeController) GetAllEmployee(c echo.Context) error {
	var emps []*dto.EmployeeResponesData
	for _, employ := range emp.repository.Get() {
		position := emp.repository.GetEmployeePositonByID(employ.PositionID)
		department := emp.repository.GetEmployeeDepartmentByID(employ.DepartmentID)

		emps = append(emps, &dto.EmployeeResponesData{
			ID:             employ.ID,
			Name:           employ.Name,
			Surname:        employ.Surname,
			PositionID:     employ.PositionID,
			DepartmentID:   employ.DepartmentID,
			PositionName:   position.Name,
			DepartmentName: department.Name,
			Salary:         employ.Salary,
			Address:        employ.Address,
			PhoneNumber:    employ.PhoneNumber,
			Email:          employ.Email,
			Gender:         employ.Gender,
			CreatedAt:      employ.CreatedAt,
			UpdatedAt:      employ.UpdatedAt,
			DeletedAt:      employ.DeletedAt,
		})
	}

	return c.JSON(http.StatusOK, dto.EmployeeRespones{
		Message:    "",
		StatusBool: true,
		Data:       emps,
	})
}

func (emp *employeeController) GetEmployeeByID(c echo.Context) error {
	id := c.Param("id")
	employ := emp.repository.GetByID(id)

	if employ.ID == 0 {
		return c.JSON(http.StatusBadRequest, dto.EmployeeRespones{
			Message:    "",
			StatusBool: false,
		})
	}

	position := emp.repository.GetEmployeePositonByID(employ.PositionID)
	department := emp.repository.GetEmployeeDepartmentByID(employ.DepartmentID)

	return c.JSON(http.StatusOK, dto.EmployeeRespones{
		Message:    "",
		StatusBool: true,
		Data: dto.EmployeeResponesData{
			ID:             employ.ID,
			Name:           employ.Name,
			Surname:        employ.Surname,
			PositionID:     employ.PositionID,
			DepartmentID:   employ.DepartmentID,
			PositionName:   position.Name,
			DepartmentName: department.Name,
			Salary:         employ.Salary,
			Address:        employ.Address,
			PhoneNumber:    employ.PhoneNumber,
			Email:          employ.Email,
			Gender:         employ.Gender,
			CreatedAt:      employ.CreatedAt,
			UpdatedAt:      employ.UpdatedAt,
			DeletedAt:      employ.DeletedAt,
		},
	})
}

func (emp *employeeController) UpdateEmployeeByID(c echo.Context) error {
	var body *dto.EmployeeRequestBody
	id := c.Param("id")
	employ := emp.repository.GetByID(id)

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, dto.EmployeeRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	if employ.ID == 0 {
		return c.JSON(http.StatusBadRequest, dto.EmployeeRespones{
			Message:    "",
			StatusBool: false,
		})
	}

	err := emp.repository.UpdateByID(id, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.EmployeeRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.EmployeeRespones{
		Message:    "update position success",
		StatusBool: true,
	})
}

func (emp *employeeController) DeleteEmployeeByID(c echo.Context) error {
	id := c.Param("id")

	employ := emp.repository.GetByID(id)
	if employ.ID == 0 {
		return c.JSON(http.StatusBadRequest, dto.EmployeeRespones{
			Message:    "",
			StatusBool: false,
		})
	}

	err := emp.repository.DeleteByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.EmployeeRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.EmployeeRespones{
		Message:    "delete position success",
		StatusBool: true,
	})
}
