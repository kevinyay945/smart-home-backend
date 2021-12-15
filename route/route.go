package route

import (
	"github.com/labstack/echo/v4"
	"smart-home-backend/model"
	v1 "smart-home-backend/route/v1"
)

type IRoute interface {
	SetRoute(g *echo.Group)
}

func NewVersion1() IRoute {
	request := model.NewRequest()
	command := model.NewCommand()
	return v1.New(v1.Version1{
		Request: request,
		Command: command,
	})
}

