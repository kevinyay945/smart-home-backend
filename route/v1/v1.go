package v1

import (
	"github.com/labstack/echo/v4"
	"smart-home-backend/model"
)

type HttpSuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type IRoute interface {
	SetRoute(g *echo.Group)
	GetCommands(ctx echo.Context) error
	CreateCommand(ctx echo.Context) error
	UpdateCommandByUUID(ctx echo.Context) error
	DeleteCommandByUUID(ctx echo.Context) error
	GetRequests(ctx echo.Context) error
	CreateRequest(ctx echo.Context) error
	UpdateRequestByUUID(ctx echo.Context) error
	DeleteRequestByUUID(ctx echo.Context) error
}

type Version1 struct {
	Request *model.MRequest
	Command *model.MCommand
}

func New(request *model.MRequest, command *model.MCommand) IRoute {
	r := new(Version1)
	r.Request = request
	r.Command = command
	return r
}

func (v *Version1) SetRoute(g *echo.Group) {
	g.GET("/command", v.GetCommands)
	g.POST("/command", v.CreateCommand)
	g.PUT("/command/:uuid", v.UpdateCommandByUUID)
	g.DELETE("/command/:uuid", v.DeleteCommandByUUID)
	g.GET("/request", v.GetRequests)
	g.POST("/request", v.CreateRequest)
	g.PUT("/request/:uuid", v.UpdateRequestByUUID)
	g.DELETE("/request/:uuid", v.DeleteRequestByUUID)
}
