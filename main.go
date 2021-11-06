package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"smart-home-backend/model"
	"smart-home-backend/route"
	"time"
)

func init() {
	for {
		fmt.Printf("PG_URL => %v \n", os.Getenv("PG_URL"))
		dsn := os.Getenv("PG_URL")
		db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if dbErr != nil {
			time.Sleep(3*time.Second)
			continue
		}
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
