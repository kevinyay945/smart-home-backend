package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"smart-home-backend/lib/pg"
	"smart-home-backend/middleware"
	"smart-home-backend/model"
	"smart-home-backend/route"
	"smart-home-backend/utils"
)

func init() {
	for {
		fmt.Printf("PG_URL => %v \n", os.Getenv("PG_URL"))
		db := pg.GetConn()
		model.Init(db)
		break
	}
}

func main() {
	e := echo.New()

	e.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: "${time_custom} [ ${method} ] STATUS: ${status}, LATENCY: ${latency_human} URI: ${uri} \n",
		Skipper: func(c echo.Context) bool {
			// dont print root request
			if c.Request().RequestURI == "/" {
				return true
			}
			// else print it out
			return false
		},
		CustomTimeFormat: "2006/01/02 15:04:05",
	}))
	e.Use(echoMiddleware.Recover())

	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	v1Route := route.NewVersion1()
	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler
	v1Route.SetRoute(e.Group("/v1"))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
