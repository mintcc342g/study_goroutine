package model

import (
	"time"

	"gorm.io/gorm"
)

// RequestBody ...
type RequestBody struct {
	UserID  uint64 `json:"user_id"`
	Name    string
	Title   string
	Content string
}

// Validate ...
func (req *RequestBody) Validate() bool {
	return req.UserID != 0 && req.Name != "" && req.Title != ""
}

// Email ...
type Email struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	UserID    uint64         `json:"user_id`
	Name      string         `json:"name"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
}

// Emails ...
type Emails []*Email
