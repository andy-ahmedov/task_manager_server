package broker

import (
	"fmt"
	"log"

	"github.com/andy-ahmedov/task_manager_server/internal/config"
	"github.com/streadway/amqp"
)

func InitRabbitMQ(cfg *config.Broker) error {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel")
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"grpc",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Failed to declare a queue")
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Failed to declare a queue")
		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s\n", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return err
}
