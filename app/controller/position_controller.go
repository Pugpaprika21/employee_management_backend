package controller

import (
	"net/http"

	"github.com/Pugpaprika21/app/dto"
	"github.com/Pugpaprika21/app/repository"
	"github.com/labstack/echo/v4"
)

type positionController struct {
	repository repository.IPositionRepository
}

func NewPositionController(repository *repository.PositionRepository) *positionController {
	return &positionController{
		repository: repository,
	}
}

func (p *positionController) CreatePosition(c echo.Context) error {
	var body *dto.PositionRequestBody

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, dto.PositionRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	if p.repository.CheckPositionNameIsAlreadyHaveIt(body) > 0 {
		return c.JSON(http.StatusBadRequest, dto.PositionRespones{
			Message:    "position name exting",
			StatusBool: false,
		})
	}

	err := p.repository.Save(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.PositionRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.PositionRespones{
		Message:    "save department success",
		StatusBool: true,
	})
}

func (p *positionController) GetAllPosition(c echo.Context) error {
	var data []*dto.PositionResponesData
	for _, pos := range p.repository.Get() {
		data = append(data, &dto.PositionResponesData{
			ID:               pos.ID,
			Name:             pos.Name,
			Description:      pos.Description,
			Salary:           pos.Salary,
			Status:           pos.Status,
			Location:         pos.Location,
			Responsibilities: pos.Responsibilities,
			CreatedAt:        pos.CreatedAt,
			UpdatedAt:        pos.UpdatedAt,
			DeletedAt:        pos.DeletedAt,
		})
	}

	return c.JSON(http.StatusOK, dto.PositionRespones{
		Message:    "",
		StatusBool: true,
		Data:       data,
	})
}

func (p *positionController) GetPositionByID(c echo.Context) error {
	id := c.Param("id")
	pos := p.repository.GetByID(id)

	if pos.ID == 0 {
		return c.JSON(http.StatusBadRequest, dto.PositionRespones{
			Message:    "",
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.PositionRespones{
		Message:    "",
		StatusBool: true,
		Data: dto.PositionResponesData{
			ID:               pos.ID,
			Name:             pos.Name,
			Description:      pos.Description,
			Salary:           pos.Salary,
			Status:           pos.Status,
			Location:         pos.Location,
			Responsibilities: pos.Responsibilities,
			CreatedAt:        pos.CreatedAt,
			UpdatedAt:        pos.UpdatedAt,
			DeletedAt:        pos.DeletedAt,
		},
	})
}

func (p *positionController) UpdatePositionByID(c echo.Context) error {
	var body *dto.PositionRequestBody
	id := c.Param("id")
	pos := p.repository.GetByID(id)

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, dto.PositionRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	if pos.ID == 0 {
		return c.JSON(http.StatusBadRequest, dto.PositionRespones{
			Message:    "",
			StatusBool: false,
		})
	}

	if p.repository.CheckPositionNameIsAlreadyHaveIt(body) > 0 {
		return c.JSON(http.StatusBadRequest, dto.PositionRespones{
			Message:    "position name exting",
			StatusBool: false,
		})
	}

	err := p.repository.UpdateByID(id, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.PositionRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.PositionRespones{
		Message:    "update position success",
		StatusBool: true,
	})
}

func (p *positionController) DeletePositionByID(c echo.Context) error {
	id := c.Param("id")

	pos := p.repository.GetByID(id)
	if pos.ID == 0 {
		return c.JSON(http.StatusBadRequest, dto.PositionRespones{
			Message:    "",
			StatusBool: false,
		})
	}

	err := p.repository.DeleteByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.PositionRespones{
			Message:    err.Error(),
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.PositionRespones{
		Message:    "delete position success",
		StatusBool: true,
	})
}
