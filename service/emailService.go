package service

import (
	"context"
	"log"
	"study_goroutine/model"
)

type emailUsecase struct {
	backgroundService BackgroundService
}

// NewEmailService ...
func NewEmailService(backgroundService BackgroundService) EmailService {
	return &emailUsecase{
		backgroundService: backgroundService,
	}
}

func (e *emailUsecase) NewEmail(ctx context.Context, req *model.RequestBody) (email *model.Email, err error) {

	// TODO: 유효성 체크

	if email, err = req.MakeEmail(); err != nil {
		log.Println("EmailService NewEmail MakeEmail Error")
		return nil, err
	}
	task, err := email.NewTask()
	if err != nil {
		log.Println("EmailService NewEmail NewTask Error")
		return nil, err
	}
	e.backgroundService.SendBackgroundTask(ctx, task)

	return email, nil
}

func (e *emailUsecase) UserEmail(ctx context.Context, userID uint64, emailID string) (email *model.Email, err error) {
	return email, nil
}

func (e *emailUsecase) UserEmails(ctx context.Context, userID uint64) (emails model.Emails, err error) {
	return emails, nil
}

func (e *emailUsecase) UpdateEmail(ctx context.Context, req *model.RequestBody) (email *model.Email, err error) {
	return email, nil
}

func (e *emailUsecase) DeleteEmail(ctx context.Context, req *model.RequestBody) (err error) {
	return nil
}
