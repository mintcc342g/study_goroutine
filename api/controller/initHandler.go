package controller

import (
	"os"
	"study_goroutine/conf"
	"study_goroutine/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	// ResponseBody ...
	ResponseBody struct {
		StatusCode int         `json:"resultCode" example:"000"`
		ResultMsg  string      `json:"resultMsg" example:"Request OK"`
		ResultData interface{} `json:"resultData,omitempty"`
	}
)

// InitHandler ...
func InitHandler(studyGoroutine *conf.ViperConfig, e *echo.Echo, signal <-chan os.Signal) error {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Default Group
	api := e.Group("/api")
	ver := api.Group("/v1")
	sys := ver.Group("/email")

	// Services
	backgroundService := service.NewBackgroundService(signal)
	emailService := service.NewEmailService(backgroundService)

	// Handlers
	newHTTPHandler(studyGoroutine, sys, emailService)

	return nil
}

func response(c echo.Context, code int, resMsg string, result ...interface{}) error {
	res := ResponseBody{
		StatusCode: code,
		ResultMsg:  resMsg,
	}
	if result != nil {
		res.ResultData = result[0]
	}

	return c.JSON(code, res)
}
