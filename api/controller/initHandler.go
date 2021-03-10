package controller

import (
	"study_goroutine/conf"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// InitHandler ...
func InitHandler(studyGoroutine *conf.ViperConfig, e *echo.Echo) error {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Default Group
	api := e.Group("/api")
	ver := api.Group("/v1")
	sys := ver.Group("/mail")

	newHTTPHandler(studyGoroutine, sys)

	return nil
}
