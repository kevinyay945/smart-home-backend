package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type HttpSuccessResponse struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
}

type Command struct {
	Uuid string `json:"uuid"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	Url string `json:"url"`
}
var allCommand []Command = []Command{{
				"bbbad37f-c6cd-47f7-907d-add1c4045559",
				time.Now(),
				time.Now(),
				"http://example.com",
			},{
			"bbbad37f-c6cd-47f7-907d-add1c4045558",
			time.Now(),
			time.Now(),
			"http://example2.com",
			}}
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/commands", func(c echo.Context) error {
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: allCommand,
		})
	})

	e.POST("/commands", func(c echo.Context) error {
		input := new(Command)
		if err := c.Bind(input); err != nil {
			return c.String(http.StatusBadRequest, "Fail to Bind Data")
		}
		_uuid, uuidErr := uuid.NewRandom()
		if uuidErr != nil {
			return c.String(http.StatusInternalServerError, "Fail to generate uuid")
		}
		input.Uuid = _uuid.String()
		input.CreateAt = time.Now()
		input.UpdateAt = time.Now()
	return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: input,
		})
	})

	e.PUT("/commands/:uuid", func(c echo.Context) error {
		input := new(Command)
		if err := c.Bind(input); err != nil {
			return c.String(http.StatusBadRequest, "Fail to Bind Data")
		}
		commandUuid := c.Param("uuid")
		output := new(Command)
		for _, command := range allCommand {
			if commandUuid == command.Uuid {
				output = &command
			}
		}
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: output,
		})
	})

	e.DELETE("/commands/:uuid", func(c echo.Context) error {
		commandUuid := c.Param("uuid")
		fmt.Println("Delete uuid", commandUuid)
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: nil,
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
