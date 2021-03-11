package service

import (
	"study_goroutine/model"
)

type emailUsecase struct {
}

// NewEmailService ...
func NewEmailService() EmailService {
	return &emailUsecase{}
}

func (e *emailUsecase) NewEmail(req *model.RequestBody) (email *model.Email, err error) {
	return &model.Email{}, nil
}

func (e *emailUsecase) UserEmail(userID uint64, emailID string) (email *model.Email, err error) {
	return email, nil
}

func (e *emailUsecase) UserEmails(userID uint64) (emails model.Emails, err error) {
	return emails, nil
}

func (e *emailUsecase) UpdateEmail(req *model.RequestBody) (email *model.Email, err error) {
	return email, nil
}

func (e *emailUsecase) DeleteEmail(req *model.RequestBody) (err error) {
	return nil
}
