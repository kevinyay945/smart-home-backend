package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"smart-home-backend/lib/pq"
	"smart-home-backend/model"
	"smart-home-backend/route"
)

func init() {
	for {
		fmt.Printf("PG_URL => %v \n", os.Getenv("PG_URL"))
		db := pq.GetConn()
		model.Init(db)
		break
	}
}

func main() {
	e := echo.New()
	v1Route := route.NewVersion1()

	v1Route.SetRoute(e.Group("/v1"))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
