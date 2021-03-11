package controller

import (
	"net/http"
	"study_goroutine/conf"
	"study_goroutine/model"
	"study_goroutine/service"

	"github.com/labstack/echo/v4"
)

// HTTPHandler ...
type HTTPHandler struct {
	studyGoroutine *conf.ViperConfig
	emailService   service.EmailService
}

func newHTTPHandler(studyGoroutine *conf.ViperConfig, eg *echo.Group, emailService service.EmailService) {
	handler := &HTTPHandler{
		studyGoroutine: studyGoroutine,
		emailService:   emailService,
	}
	eg.POST("/send", handler.SendEmail)
	eg.GET("/:emailID", handler.Email)
}

// SendEmail ...
func (handler *HTTPHandler) SendEmail(c echo.Context) (err error) {

	req := &model.RequestBody{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if !req.Validate() {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	email, err := handler.emailService.NewEmail(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, email)
}

// Email ...
func (handler *HTTPHandler) Email(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Hello, World!") // Test -> http://localhost:1323/api/v1/mail/1
}
