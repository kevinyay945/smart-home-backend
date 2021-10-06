package main

import (
	"github.com/google/uuid"
	"net/http"
	"smart-home-backend/model"
	"time"

	"github.com/labstack/echo"
)

type HttpSuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/commands", func(c echo.Context) error {
		var output []model.Command
		command := model.NewCommand()
		output = command.Get()
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: struct {
				Commands []model.Command `json:"commands"`
			}{
				Commands: output,
			},
		})
	})

	e.POST("/commands", func(c echo.Context) error {
		command := model.NewCommand()
		input := new(model.Command)
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
		result := command.Save(input)
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: struct {
				Command model.Command `json:"command"`
			}{Command: result},
		})
	})

	e.PUT("/commands/:uuid", func(c echo.Context) error {
		command := model.NewCommand()
		input := new(model.Command)
		if err := c.Bind(input); err != nil {
			return c.String(http.StatusBadRequest, "Fail to Bind Data")
		}
		commandUuid := c.Param("uuid")
		result := command.UpdateOne(commandUuid, command)
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: struct {
				Command model.Command `json:"command"`
			}{
				Command: result,
			},
		})
	})

	e.DELETE("/commands/:uuid", func(c echo.Context) error {
		command := model.NewCommand()
		commandUuid := c.Param("uuid")
		_ = command.Delete(commandUuid)
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data:   nil,
		})
	})

	e.GET("/requests", func(c echo.Context) error {
		request := model.NewRequest()
		result := request.Get()
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: struct {
				Request []model.Request `json:"requests"`
			}{
				Request: result,
			},
		})
	})

	e.POST("/requests", func(c echo.Context) error {
		request := model.NewRequest()
		input := new(model.Request)
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
		result := request.Save(input)
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: struct {
				Request model.Request `json:"request"`
			}{Request: result},
		})
	})

	e.PUT("/requests/:uuid", func(c echo.Context) error {
		request := model.NewRequest()
		input := new(model.Request)
		if err := c.Bind(input); err != nil {
			return c.String(http.StatusBadRequest, "Fail to Bind Data")
		}
		requestUuid := c.Param("uuid")
		result := request.Update(requestUuid, request)
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data: struct {
				Request model.Request `json:"request"`
			}{
				Request: result,
			},
		})
	})

	e.DELETE("/requests/:uuid", func(c echo.Context) error {
		request := model.NewRequest()
		commandUuid := c.Param("uuid")
		_ = request.Delete(commandUuid)
		return c.JSON(http.StatusOK, HttpSuccessResponse{
			Status: "success",
			Data:   nil,
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
