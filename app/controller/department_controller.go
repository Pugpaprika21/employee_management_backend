package controller

import (
	"net/http"

	"github.com/Pugpaprika21/app/dto"
	"github.com/Pugpaprika21/app/repository"
	"github.com/labstack/echo/v4"
)

type departmentController struct {
	repository repository.IDepartmentRepository
}

func NewDepartmentController(repository *repository.DepartmentRepository) *departmentController {
	return &departmentController{
		repository: repository,
	}
}

func (d *departmentController) CreateDepartment(c echo.Context) error {
	var body *dto.DepartmentRequestBody

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, dto.DepartmentRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	if d.repository.CheckDepartmentNameIsAlreadyHaveIt(body) > 0 {
		return c.JSON(http.StatusBadRequest, dto.DepartmentRespones{
			Message:    "department name exting",
			StatusBool: false,
		})
	}

	err := d.repository.Save(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.DepartmentRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.DepartmentRespones{
		Message:    "save department success",
		StatusBool: true,
	})
}

func (d *departmentController) GetAllDepartment(c echo.Context) error {
	var data []*dto.DepartmentResponesData
	for _, dpm := range d.repository.Get() {
		data = append(data, &dto.DepartmentResponesData{
			ID:        dpm.ID,
			Name:      dpm.Name,
			Location:  dpm.Location,
			CreatedAt: dpm.CreatedAt,
			UpdatedAt: dpm.UpdatedAt,
			DeletedAt: dpm.DeletedAt,
		})
	}

	return c.JSON(http.StatusOK, dto.DepartmentRespones{
		Message:    "",
		StatusBool: true,
		Data:       data,
	})
}

func (d *departmentController) GetDepartmentByID(c echo.Context) error {
	id := c.Param("id")
	dep := d.repository.GetByID(id)

	if dep.ID == 0 {
		return c.JSON(http.StatusBadRequest, dto.DepartmentRespones{
			Message:    "",
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.DepartmentRespones{
		Message:    "",
		StatusBool: true,
		Data: dto.DepartmentResponesData{
			ID:        dep.ID,
			Name:      dep.Name,
			Location:  dep.Location,
			CreatedAt: dep.CreatedAt,
			UpdatedAt: dep.UpdatedAt,
			DeletedAt: dep.DeletedAt,
		},
	})
}

func (d *departmentController) UpdateDepartmentByID(c echo.Context) error {
	var body *dto.DepartmentRequestBody
	id := c.Param("id")
	dep := d.repository.GetByID(id)

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, dto.DepartmentRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	if dep.ID == 0 {
		return c.JSON(http.StatusBadRequest, dto.DepartmentRespones{
			Message:    "",
			StatusBool: false,
		})
	}

	if d.repository.CheckDepartmentNameIsAlreadyHaveIt(body) > 0 {
		return c.JSON(http.StatusBadRequest, dto.DepartmentRespones{
			Message:    "department name exting",
			StatusBool: false,
		})
	}

	err := d.repository.UpdateByID(id, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.DepartmentRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.DepartmentRespones{
		Message:    "update department success",
		StatusBool: true,
	})
}

func (d *departmentController) DeleteDepartmentByID(c echo.Context) error {
	id := c.Param("id")

	dep := d.repository.GetByID(id)
	if dep.ID == 0 {
		return c.JSON(http.StatusBadRequest, dto.DepartmentRespones{
			Message:    "",
			StatusBool: false,
		})
	}

	err := d.repository.DeleteByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.DepartmentRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.DepartmentRespones{
		Message:    "delete department success",
		StatusBool: true,
	})
}
