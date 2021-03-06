package v1

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"smart-home-backend/lib/pg/schema"
	"smart-home-backend/model"
	"time"
)

var (
	newUuidV4 = uuid.NewRandom
	getNow    = time.Now
)

type requestRoute struct {
	Request model.IRequest
}

func (r *requestRoute) GetRequests(ctx echo.Context) error {
	result, getErr := r.Request.Get()
	if getErr != nil {
		fmt.Println("Get Requests Error")
		return echo.NewHTTPError(http.StatusInternalServerError, getErr)
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

func (r *requestRoute) CreateRequest(ctx echo.Context) error {
	input := new(schema.Request)
	if err := ctx.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	_uuid, uuidErr := newUuidV4()
	if uuidErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Fail to generate uuid")
	}
	input.Uuid = _uuid.String()
	input.CreateAt = getNow()
	input.UpdateAt = getNow()
	if err := ctx.Validate(input); err != nil {
		return err
	}
	result, _ := r.Request.Save(input)
	return ctx.JSON(http.StatusCreated, HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Request schema.Request `json:"request"`
		}{
			Request: result,
		},
	})
}

func (r *requestRoute) UpdateRequestByUUID(ctx echo.Context) error {
	input := new(schema.Request)
	if err := ctx.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	requestUuid := ctx.Param("uuid")
	result, _ := r.Request.Update(requestUuid, input)
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Request schema.Request `json:"request"`
		}{
			Request: result,
		},
	})
}

func (r *requestRoute) DeleteRequestByUUID(ctx echo.Context) error {
	commandUuid := ctx.Param("uuid")
	err := r.Request.Delete(commandUuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, HttpSuccessResponse{
		Status: "success",
		Data:   nil,
	})
}
