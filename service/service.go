package service

import (
	"study_goroutine/model"
)

// EmailService ...
type EmailService interface {
	NewEmail(req *model.RequestBody) (email *model.Email, err error)
	UserEmail(userID uint64, emailID string) (email *model.Email, err error)
	UserEmails(userID uint64) (emails model.Emails, err error)
	UpdateEmail(req *model.RequestBody) (email *model.Email, err error)
	DeleteEmail(req *model.RequestBody) (err error)
}
