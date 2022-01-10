package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"smart-home-backend/lib/pg"
	"smart-home-backend/middleware"
	"smart-home-backend/model"
	"smart-home-backend/route"
)

func init() {
	for {
		fmt.Printf("PG_URL => %v \n", os.Getenv("PG_URL"))
		db := pg.GetConn()
		model.Init(db)
		break
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	v1Route := route.NewVersion1()
	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler
	v1Route.SetRoute(e.Group("/v1"))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
