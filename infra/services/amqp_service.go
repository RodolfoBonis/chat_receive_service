package services

import (
	"github.com/streadway/amqp"
)

type AmqpService struct {
	UrlConnection string
	QueueName     string
}

func (amqpService *AmqpService) OpenAmqpConnection() *amqp.Channel {
	amqpServerURL := amqpService.UrlConnection

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}

	_, err = channelRabbitMQ.QueueDeclare(
		amqpService.QueueName, // queue name
		true,                  // durable
		false,                 // auto delete
		false,                 // exclusive
		false,                 // no wait
		nil,                   // arguments
	)

	if err != nil {
		panic(err)
	}

	return channelRabbitMQ
}
