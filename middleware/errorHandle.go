package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type DefaultHttpFailResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	if err := c.JSON(code, DefaultHttpFailResponse{
		Status:  "fail",
		Message: err.Error(),
	}); err != nil {
		c.Logger().Error(err)
	}
}
