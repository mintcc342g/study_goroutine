package controller

import (
	"study_goroutine/conf"
	"study_goroutine/service"

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
	sys := ver.Group("/email")

	// Services
	emailService := service.NewEmailService()

	// Handlers
	newHTTPHandler(studyGoroutine, sys, emailService)

	return nil
}
