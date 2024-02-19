package rabbitmq

import (
	"log"

	"github.com/andy-ahmedov/task_manager_server/internal/config"
	"github.com/sirupsen/logrus"
)

func InitRabbitMQ(cfg *config.Broker, logg *logrus.Logger) {
	conn, err := ConnectToTCP(cfg, logg)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := CreateChannel(conn, logg)
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := DeclareQueue(ch, "LogItemsQueue", logg)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ConsumeChannel(q, ch, logg)
	if err != nil {
		log.Fatal(err)
	}

	go QueueProcessing(msgs, logg)

	select {}
}
