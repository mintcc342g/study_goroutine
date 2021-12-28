package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"study-goroutine/model"
	"study-goroutine/repository"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type backgroundUsecase struct {
	mqRepo    repository.MQRepository
	mqChannel <-chan amqp.Delivery
}

// NewBackgroundService ...
func NewBackgroundService(mqRepo repository.MQRepository, cancel <-chan os.Signal) BackgroundService {
	bg := &backgroundUsecase{
		mqRepo:    mqRepo,
		mqChannel: mqRepo.DeliveryMessage(),
	}
	go bg.runBackgroundProcess(cancel)
	return bg
}

func (bg *backgroundUsecase) SendBackgroundTask(ctx context.Context, task *model.BackgroundTask) error {
	rawTask, err := json.Marshal(task)
	if err != nil {
		log.Println("backgroundService SendBackgroundTask Marshal Error", err)
		return err
	}
	bg.mqRepo.PublishMessage(rawTask)

	return nil
}

func (bg *backgroundUsecase) runBackgroundProcess(cancel <-chan os.Signal) {
	ctx := context.Background()

	defer func() { // 고루틴 충돌 시의 panic recover
		if v := recover(); v != nil {
			go bg.runBackgroundProcess(cancel) // 고루틴 재시작
		}
	}()

	select {
	case <-cancel:
		log.Println("shutdown background process...")
		return
	default:
		if err := bg.consumeMessage(ctx); err != nil {
			log.Println("backgroundService runBackgroundProcess consumeMessage Error", err)
		}
	}
}

func (bg *backgroundUsecase) consumeMessage(ctx context.Context) error {
	for d := range bg.mqChannel {
		var task *model.BackgroundTask
		if err := json.Unmarshal(d.Body, &task); err != nil {
			d.Nack(false, true)
			log.Println("backgroundService consumeMessage Unmarshal Error", err)
			return err
		}

		if err := bg.executeTask(ctx, task); err != nil {
			d.Nack(false, true) // 에러 발생 시, 메세지가 소비되지 않고 다시 mq로 들어감. 그래서 for 가 계속 돌게 됨.
			log.Println("backgroundService consumeMessage executeTask Error", err)
			continue
		}

		d.Ack(false)
		log.Println("backgroundService consumeMessage Success...")
	}

	return nil
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

	if err := email.NewError(); err != nil { // runBackgroundProcess() 에서 d.Nack() 후 continue를 확인하기 위해 에러 발생시킴.
		return err
	}

	// TODO: 이메일 데이터 DB 저장

	log.Println("backgroundService Successs to send the email...")
	return nil
}
