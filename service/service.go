package service

import (
	"context"
	"study_goroutine/model"
)

// EmailService ...
type EmailService interface {
	NewEmail(ctx context.Context, req *model.RequestBody) (email *model.Email, err error)
	UserEmail(ctx context.Context, userID uint64, emailID string) (email *model.Email, err error)
	UserEmails(ctx context.Context, userID uint64) (emails model.Emails, err error)
	UpdateEmail(ctx context.Context, req *model.RequestBody) (email *model.Email, err error)
	DeleteEmail(ctx context.Context, req *model.RequestBody) (err error)
}

// BackgroundService ...
type BackgroundService interface {
	SendBackgroundTask(ctx context.Context, task *model.BackgroundTask)
}
