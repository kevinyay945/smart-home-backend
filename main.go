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
			Data: struct {
				Commands []Command `json:"commands"`
			}{
				Commands: allCommand,
			},
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
			Data: struct {
				Command Command `json:"command"`
			}{Command: *input},
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
			Data: struct {
				Command Command `json:"command"`
			}{
				Command: *output,
			},
		})
	})

	e.DELETE("/commands/:uuid", func(c echo.Context) error {
		requestUuid := c.Param("uuid")
		fmt.Println("Delete uuid", requestUuid)
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: nil,
		})
	})

	type Request struct {
		Uuid string `json:"uuid"`
		CreateAt time.Time `json:"createAt"`
		UpdateAt time.Time `json:"updateAt"`
		Name string `json:"name"`
	}

	var allRequest []Request = []Request{{
		"324e672e-3512-41b0-97dc-065c334f8f7a",
		time.Now(),
		time.Now(),
		"Request 1",
	},{
		"6737a85c-3ea4-48ed-bfed-4bdff74b189b",
		time.Now(),
		time.Now(),
		"Request 2",
	}}

	e.GET("/requests", func(c echo.Context) error {
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: struct {
				Request []Request `json:"requests"`
			}{
				Request: allRequest,
			},
		})
	})

	e.POST("/requests", func(c echo.Context) error {
		input := new(Request)
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
			Data: struct {
				Request Request `json:"request"`
			}{Request: *input},
		})
	})

	e.PUT("/requests/:uuid", func(c echo.Context) error {
		input := new(Request)
		if err := c.Bind(input); err != nil {
			return c.String(http.StatusBadRequest, "Fail to Bind Data")
		}
		requestUuid := c.Param("uuid")
		output := new(Request)
		for _, request := range allRequest {
			if requestUuid == request.Uuid {
				output.Uuid = request.Uuid
				output.CreateAt = request.CreateAt
				output.Name = input.Name
				output.UpdateAt = time.Now()
			}
		}
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: struct {
				Request Request `json:"request"`
			}{
				Request: *output,
			},
		})
	})

	e.DELETE("/requests/:uuid", func(c echo.Context) error {
		commandUuid := c.Param("uuid")
		fmt.Println("Delete uuid", commandUuid)
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: nil,
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
