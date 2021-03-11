package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"study_goroutine/model"
)

type backgroundUsecase struct {
	taskChannel chan *model.BackgroundTask
}

// NewBackgroundService ...
func NewBackgroundService(cancel <-chan os.Signal) BackgroundService {
	bg := &backgroundUsecase{
		taskChannel: make(chan *model.BackgroundTask, 100),
	}
	go bg.runBackgroundProcess(cancel)
	return bg
}

func (bg *backgroundUsecase) SendBackgroundTask(ctx context.Context, task *model.BackgroundTask) {
	bg.taskChannel <- task
	return
}

func (bg *backgroundUsecase) runBackgroundProcess(cancel <-chan os.Signal) {
	ctx := context.Background()

	defer func() {
		if v := recover(); v != nil {
			go bg.runBackgroundProcess(cancel)
		}
	}()

	for {
		select {
		case <-cancel:
			log.Println("shutdown background process...")
			return
		case task := <-bg.taskChannel:
			bg.executeTask(ctx, task)
		}
	}
}

func (bg *backgroundUsecase) executeTask(ctx context.Context, task *model.BackgroundTask) {

	switch task.TaskType {
	case model.TaskEventType(model.TaskEventTypeEmailSend):
		err := bg.SendEmail(ctx, task.TaskData)
		if err != nil {
			log.Println("BackgroundService executeTask Send Email Error")
		}

	default:
		log.Println("invalid task type")
		return
	}

	return
}

func (bg *backgroundUsecase) SendEmail(ctx context.Context, rawEmail []byte) error {

	// 비동기 처리 확인을 위한 sleep
	fmt.Println("backgroundService Ready for sending the email...") // log로 바꿀 경우, response가 먼저 나가고 log가 나중에 뜨기 시작함.
	time.Sleep(time.Second * 2)
	fmt.Println("backgroundService Start sending the email...")

	var email *model.Email
	if err := json.Unmarshal(rawEmail, &email); err != nil {
		log.Println("BackgroundService SendEmail Unmarshal Error")
		return err
	}

	// TODO: 이메일 데이터 DB 저장

	log.Println("backgroundService Successs to send the email...")
	return nil
}
