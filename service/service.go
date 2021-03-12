package service

import (
	"context"
	"study_goroutine/model"
)

// EmailService ...
type EmailService interface {
	NewEmail(ctx context.Context, req *model.RequestBody) (*model.Email, error)
	UserEmail(ctx context.Context, userID uint64, emailID string) (*model.Email, error)
	UserEmails(ctx context.Context, userID uint64) (model.Emails, error)
	UpdateEmail(ctx context.Context, req *model.RequestBody) (*model.Email, error)
	DeleteEmail(ctx context.Context, req *model.RequestBody) error
}

// BackgroundService ...
type BackgroundService interface {
	SendBackgroundTask(ctx context.Context, task *model.BackgroundTask) error
}
