package controller

import (
	"net/http"
	"study_goroutine/conf"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitHandler ...
func InitHandler(studyGoroutine *conf.ViperConfig, e *echo.Echo) error {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	return nil
}
