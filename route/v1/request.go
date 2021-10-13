package v1

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"smart-home-backend/model"
	"time"
)

func (v *Version1) GetRequests(ctx echo.Context) error {
	request := model.NewRequest()
	result := request.Get()
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Request []model.Request `json:"requests"`
		}{
			Request: result,
		},
	})
}

func (v *Version1) CreateRequest(ctx echo.Context) error {
	request := model.NewRequest()
	input := new(model.Request)
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
	result := request.Save(input)
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Request model.Request `json:"request"`
		}{Request: result},
	})
}

func (v *Version1) UpdateRequestByUUID(ctx echo.Context) error {
	request := model.NewRequest()
	input := new(model.Request)
	if err := ctx.Bind(input); err != nil {
		return ctx.String(http.StatusBadRequest, "Fail to Bind Data")
	}
	requestUuid := ctx.Param("uuid")
	result := request.Update(requestUuid, request)
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Request model.Request `json:"request"`
		}{
			Request: result,
		},
	})
}

func (v *Version1) DeleteRequestByUUID(ctx echo.Context) error {
	request := model.NewRequest()
	commandUuid := ctx.Param("uuid")
	_ = request.Delete(commandUuid)
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data:   nil,
	})
}
