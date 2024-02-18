package rabbitmq

import (
	"fmt"
	"log"

	"github.com/andy-ahmedov/task_manager_server/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
	// "github.com/streadway/amqp"
)

func InitRabbitMQ(cfg *config.Broker) {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
		// return err
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel")
		// return err
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"LogItemsQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Failed to declare a queue")
		// return err
		log.Fatal(err)
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
		// return err
		log.Fatal(err)
	}

	// var forever chan struct{}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	// go func() {
	for d := range msgs {
		log.Printf("Received a message: %s\n", d.Body)
	}
	// }()

	// <-forever

	// return err
}
