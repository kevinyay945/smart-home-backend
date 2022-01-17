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
	command := g.Group("/commands")
	{
		route := &commandRoute{
			Command: v.Command,
		}
		command.GET("", route.GetCommands)
		command.POST("", route.CreateCommand)
		command.PUT("/:uuid", route.UpdateCommandByUUID)
		command.DELETE("/:uuid", route.DeleteCommandByUUID)
	}
	request := g.Group("/requests")
	{
		route := &requestRoute{
			Request: v.Request,
		}
		request.GET("", route.GetRequests)
		request.POST("", route.CreateRequest)
		request.PUT("/:uuid", route.UpdateRequestByUUID)
		request.DELETE("/:uuid", route.DeleteRequestByUUID)
	}
}
