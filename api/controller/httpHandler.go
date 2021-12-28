package controller

import (
	"context"
	"log"
	"net/http"
	"study-goroutine/conf"
	"study-goroutine/model"
	"study-goroutine/service"

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
	eg.GET("/:id", handler.Email)
}

// SendEmail ...
func (handler *HTTPHandler) SendEmail(c echo.Context) (err error) {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	req := &model.RequestBody{}
	if err := c.Bind(req); err != nil {
		log.Println("HTTPHandler SendEmail Bind Error")
		return response(c, http.StatusBadRequest, "invalid request")
	}
	if !req.Validate() {
		return response(c, http.StatusBadRequest, "invalid request")
	}

	email, err := handler.emailService.NewEmail(ctx, req)
	if err != nil {
		log.Println("HTTPHandler NewEmail Error")
		return response(c, http.StatusNotAcceptable, "err", err) // TODO: status code 수정
	}

	return response(c, http.StatusAccepted, "Send Email OK", email)
}

// Email ...
func (handler *HTTPHandler) Email(c echo.Context) (err error) {
	return response(c, http.StatusOK, "Get Email OK", "Hello, World!") // Test -> http://localhost:1323/api/v1/mail/1
}
