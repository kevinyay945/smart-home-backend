package v1

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"smart-home-backend/lib/pg/schema"
	"smart-home-backend/model"
	"time"
)

type commandRoute struct {
	Command model.ICommand
}

func (r *commandRoute) GetCommands(ctx echo.Context) error {
	var output []schema.Command
	output, getErr := r.Command.Get()
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr)
	}
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Commands []schema.Command `json:"commands"`
		}{
			Commands: output,
		},
	})
}

func (r *commandRoute) CreateCommand(ctx echo.Context) error {
	input := new(schema.Command)
	if err := ctx.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	_uuid, uuidErr := uuid.NewRandom()
	if uuidErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, uuidErr)
	}
	input.Uuid = _uuid.String()
	input.CreateAt = time.Now()
	input.UpdateAt = time.Now()
	if err := ctx.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	result, saveErr := r.Command.Save(input)
	if saveErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, saveErr)
	}
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Command schema.Command `json:"command"`
		}{Command: result},
	})
}

func (r *commandRoute) UpdateCommandByUUID(ctx echo.Context) error {
	input := new(schema.Command)
	if err := ctx.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	commandUuid := ctx.Param("uuid")
	result, updateErr := r.Command.UpdateOne(commandUuid, input)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr)
	}

	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Command schema.Command `json:"command"`
		}{
			Command: result,
		},
	})
}

func (r *commandRoute) DeleteCommandByUUID(ctx echo.Context) error {
	commandUuid := ctx.Param("uuid")
	err := r.Command.Delete(commandUuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data:   nil,
	})
}
