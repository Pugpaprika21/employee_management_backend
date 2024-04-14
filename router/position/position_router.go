package position

import (
	"github.com/Pugpaprika21/app/controller"
	"github.com/Pugpaprika21/app/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type positionRouter struct {
	controller controller.IPositionController
}

func NewPositionRouter() *positionRouter {
	return &positionRouter{}
}

func (p *positionRouter) Register(e *echo.Echo, db *gorm.DB) {
	repository := repository.NewPositionRepository(db)
	p.controller = controller.NewPositionController(repository)

	g := e.Group("/api/v1")
	g.POST("/createPosition", p.controller.CreatePosition).Name = "createPosition"
	g.GET("/getAllPosition", p.controller.GetAllPosition).Name = "getAllPosition"
	g.GET("/getPositionById/:id", p.controller.GetPositionByID).Name = "getPositionById"
	g.PUT("/updatePositionById/:id", p.controller.UpdatePositionByID).Name = "updatePositionById"
	g.DELETE("/deletePositionById/:id", p.controller.DeletePositionByID).Name = "deletePositionById"
}
