package controller

import (
	"net/http"
	"study_goroutine/conf"

	"github.com/labstack/echo/v4"
)

// HTTPHandler ...
type HTTPHandler struct {
	studyGoroutine *conf.ViperConfig
}

func newHTTPHandler(studyGoroutine *conf.ViperConfig, eg *echo.Group) {
	handler := &HTTPHandler{
		studyGoroutine: studyGoroutine,
	}
	eg.GET("/:userID", handler.UserEmail)
}

// UserEmail ...
func (h *HTTPHandler) UserEmail(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Hello, World!") // Test -> http://localhost:1323/api/v1/mail/1
}
