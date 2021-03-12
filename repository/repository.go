package repository

import "github.com/streadway/amqp"

// MQRepository ...
type MQRepository interface {
	PublishMessage(body []byte)
	DeliveryMessage() <-chan amqp.Delivery
}
