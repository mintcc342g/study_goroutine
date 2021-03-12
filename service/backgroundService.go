package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"study_goroutine/model"

	"github.com/pkg/errors"
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

	defer func() { // 고루틴 충돌 시의 panic recover
		if v := recover(); v != nil {
			go bg.runBackgroundProcess(cancel) // 고루틴 재시작
		}
	}()

	for { // exectueTask 에서 실패가 났을 경우, 다시 시도하기 위한 무한 루프
		select {
		case <-cancel:
			log.Println("shutdown background process...")
			return
		case task := <-bg.taskChannel:
			if err := bg.executeTask(ctx, task); err != nil { // TODO: 에러 종류에 따라 처리 방식 나누기
				log.Println(err)
				bg.taskChannel <- task
			}
		}
	}
}

func (bg *backgroundUsecase) executeTask(ctx context.Context, task *model.BackgroundTask) error {

	switch task.TaskType {
	case model.TaskEventType(model.TaskEventTypeEmailSend):
		if err := bg.sendEmail(ctx, task.TaskData); err != nil {
			log.Println("BackgroundService executeTask Send Email Error", err)
			return err
		}

	default:
		log.Println("invalid task type")
		return errors.Errorf("ivalid task type") // TODO: wrap errors
	}

	return nil
}

func (bg *backgroundUsecase) sendEmail(ctx context.Context, rawEmail []byte) error {

	// 비동기 처리 확인을 위한 sleep과 print
	fmt.Println("backgroundService Ready for sending the email...") // log로 바꿀 경우, response가 먼저 나가고 log가 나중에 뜨기 시작함.
	time.Sleep(time.Second * 2)
	fmt.Println("backgroundService Start sending the email...")

	var email *model.Email
	if err := json.Unmarshal(rawEmail, &email); err != nil {
		log.Println("BackgroundService SendEmail Unmarshal Error")
		return err
	}

	log.Println("Got Email", *email)

	if err := email.NewError(); err != nil { // runBackgroundProcess() 에서 무한 루프를 확인하기 위해 에러 발생시킴.
		return err
	}

	// TODO: 이메일 데이터 DB 저장

	log.Println("backgroundService Successs to send the email...")
	return nil
}
