package v1

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"smart-home-backend/lib/pq/schema"
	"smart-home-backend/model"
	"time"
)

func (v *Version1) GetCommands(ctx echo.Context) error {
	var output []schema.Command
	command := model.NewCommand()
	output, _ = command.Get()
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
		return ctx.String(http.StatusBadRequest, "Fail to Bind Data")
	}
	_uuid, uuidErr := uuid.NewRandom()
	if uuidErr != nil {
		return ctx.String(http.StatusInternalServerError, "Fail to generate uuid")
	}
	input.Uuid = _uuid.String()
	input.CreateAt = time.Now()
	input.UpdateAt = time.Now()
	result, _ := command.Save(input)
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
		return ctx.String(http.StatusBadRequest, "Fail to Bind Data")
	}
	commandUuid := ctx.Param("uuid")
	result, _ := command.UpdateOne(commandUuid, input)
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
	_, _ = command.Delete(commandUuid)
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data:   nil,
	})
}

