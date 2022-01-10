package v1

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"smart-home-backend/lib/pg/schema"
	"smart-home-backend/model"
	"time"
)

func (v *Version1) GetRequests(ctx echo.Context) error {
	request := model.NewRequest()
	result, getErr := request.Get()
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr)
	}
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Request []schema.Request `json:"requests"`
		}{
			Request: result,
		},
	})
}

func (v *Version1) CreateRequest(ctx echo.Context) error {
	request := model.NewRequest()
	input := new(schema.Request)
	if err := ctx.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	_uuid, uuidErr := uuid.NewRandom()
	if uuidErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Fail to generate uuid")
	}
	input.Uuid = _uuid.String()
	input.CreateAt = time.Now()
	input.UpdateAt = time.Now()
	if err := ctx.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	result, _ := request.Save(input)
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Request schema.Request `json:"request"`
		}{Request: result},
	})
}

func (v *Version1) UpdateRequestByUUID(ctx echo.Context) error {
	request := model.NewRequest()
	input := new(schema.Request)
	if err := ctx.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	requestUuid := ctx.Param("uuid")
	result, _ := request.Update(requestUuid, input)
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Request schema.Request `json:"request"`
		}{
			Request: result,
		},
	})
}

func (v *Version1) DeleteRequestByUUID(ctx echo.Context) error {
	request := model.NewRequest()
	commandUuid := ctx.Param("uuid")
	err := request.Delete(commandUuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data:   nil,
	})
}
