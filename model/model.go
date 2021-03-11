package model

// TaskEventType ...
type TaskEventType int

// TaskEventType List ...
const (
	TaskEventTypeNone int = iota
	TaskEventTypeEmailSend
)

// BackgroundTask ...
type BackgroundTask struct {
	TaskData []byte
	TaskType TaskEventType
}
