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
}

type Version1 struct {
	Request model.IRequest
	Command model.ICommand
}

func New(input Version1) IRoute {
	r := new(Version1)
	r.Request = input.Request
	r.Command = input.Command
	return r
}

func (v *Version1) SetRoute(g *echo.Group) {
	g.GET("/commands", v.GetCommands)
	g.POST("/commands", v.CreateCommand)
	g.PUT("/commands/:uuid", v.UpdateCommandByUUID)
	g.DELETE("/commands/:uuid", v.DeleteCommandByUUID)
	g.GET("/requests", v.GetRequests)
	g.POST("/requests", v.CreateRequest)
	g.PUT("/requests/:uuid", v.UpdateRequestByUUID)
	g.DELETE("/requests/:uuid", v.DeleteRequestByUUID)
}
