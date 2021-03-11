package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// RequestBody ...
type RequestBody struct {
	SenderAddress   string
	ReceiverAddress string
	Title           string
	Content         string
}

// Validate ...
func (req *RequestBody) Validate() bool {
	return req.SenderAddress != "" && req.ReceiverAddress != "" && req.Title != ""
}

// MakeEmail ...
func (req *RequestBody) MakeEmail() (*Email, error) {

	// TODO: 이메일 주소 regex

	return &Email{
		SenderAddress:   req.SenderAddress,
		ReceiverAddress: req.ReceiverAddress,
		Title:           req.Title,
		Content:         req.Content,
	}, nil
}

// Email ...
type Email struct {
	ID              uint64         `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt       *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt       *time.Time     `json:"updatedAt,omitempty"`
	DeletedAt       gorm.DeletedAt `json:"deletedAt,omitempty"`
	SenderAddress   string         `json:"senderAddress"`
	ReceiverAddress string         `json:"receiverAddress"`
	Title           string         `json:"title"`
	Content         string         `json:"content"`
}

// NewTask ...
func (e *Email) NewTask() (*BackgroundTask, error) {
	byt, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	return &BackgroundTask{
		TaskData: byt,
		TaskType: TaskEventType(TaskEventTypeEmailSend),
	}, nil
}

// Emails ...
type Emails []*Email
