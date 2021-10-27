package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"smart-home-backend/model"
	"smart-home-backend/route"
)

func init() {
	for {
		dsn := "postgres://dsjsenja:tMheNWxbf3b5U7xL65XjKiouKTP_1CXd@satao.db.elephantsql.com/dsjsenja"
		db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if dbErr != nil {
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
