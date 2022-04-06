package repositories

import (
	"log"

	"chat_receive_service/infra/services"
	"chat_receive_service/models"
	"chat_receive_service/utils"
)

type MessageRepository struct{}

func (m *MessageRepository) SendMessage(message models.MessageModel) {

}

func (m *MessageRepository) ListenQueueMessages() {
	amqpService := services.AmqpService{
		UrlConnection: utils.GetEnv("AMQP_URL", "amqp://guest:guest@localhost:5672/"),
		QueueName:     "chat_messages",
	}

	channel := amqpService.OpenAmqpConnection()

	channel.Qos(10, 0, false)

	msgs, err := channel.Consume(
		"chat_messages",
		"/",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
