package repository

import (
	"log"

	"github.com/streadway/amqp"
)

type mqUsecase struct {
	mqChannel    *amqp.Channel
	mqRoutingKey string
}

// NewMQRepository ...
func NewMQRepository(mqCh *amqp.Channel, mqRoutingKey string) MQRepository {
	return &mqUsecase{
		mqChannel:    mqCh,
		mqRoutingKey: mqRoutingKey,
	}
}

func (repo *mqUsecase) PublishMessage(body []byte) {
	err := repo.mqChannel.Publish(
		"",                // exchange
		repo.mqRoutingKey, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		log.Println("Failed to publish a message", err)
	}
	log.Printf(" [x] Sent %s", body)
}

func (repo *mqUsecase) DeliveryMessage() <-chan amqp.Delivery {
	msgs, err := repo.mqChannel.Consume(
		repo.mqRoutingKey, // queue
		"",                // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		log.Println("Failed to register a consumer", err)
	}

	return msgs
}
