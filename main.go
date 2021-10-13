package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"smart-home-backend/route"
)


func main() {
	e := echo.New()

	v1Route := route.NewVersion1()

	v1Route.SetRoute(e.Group("/v1"))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
