package v1

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"smart-home-backend/lib/pg/schema"
	"smart-home-backend/model"
	"time"
)

func (v *Version1) GetCommands(ctx echo.Context) error {
	var output []schema.Command
	command := model.NewCommand()
	output, getErr := command.Get()
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

func (v *Version1) CreateCommand(ctx echo.Context) error {
	command := model.NewCommand()
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
	result, saveErr := command.Save(input)
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

func (v *Version1) UpdateCommandByUUID(ctx echo.Context) error {
	command := model.NewCommand()
	input := new(schema.Command)
	if err := ctx.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	commandUuid := ctx.Param("uuid")
	result, updateErr := command.UpdateOne(commandUuid, input)
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

func (v *Version1) DeleteCommandByUUID(ctx echo.Context) error {
	command := model.NewCommand()
	commandUuid := ctx.Param("uuid")
	err := command.Delete(commandUuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data:   nil,
	})
}
