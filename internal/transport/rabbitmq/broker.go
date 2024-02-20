package rabbitmq

import (
	"context"
	"log"

	"github.com/andy-ahmedov/task_manager_server/internal/config"
	"github.com/andy-ahmedov/task_manager_server/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type LogItemService interface {
	Create(ctx context.Context, logItem domain.LogItem) error
}

type Broker struct {
	Conn    *amqp.Connection
	Ch      *amqp.Channel
	Q       *amqp.Queue
	Log     *logrus.Logger
	Service LogItemService
}

func NewBroker(cfg *config.Broker, logg *logrus.Logger, logItemService LogItemService) *Broker {
	conn, err := ConnectToTCP(cfg, logg)
	if err != nil {
		log.Fatal(err)
	}

	ch, err := CreateChannel(conn, logg)
	if err != nil {
		log.Fatal(err)
	}

	q, err := DeclareQueue(ch, "LogItemsQueue", logg)
	if err != nil {
		log.Fatal(err)
	}

	return &Broker{
		Conn:    conn,
		Ch:      ch,
		Q:       &q,
		Log:     logg,
		Service: logItemService,
	}
}

func (b *Broker) RunBrokerServer(ctx context.Context) {
	msgs, err := ConsumeChannel(*b.Q, b.Ch, b.Log)
	if err != nil {
		b.Log.Fatal(err)
	}

	go b.QueueProcessing(ctx, msgs, b.Log)
	defer b.Conn.Close()
	defer b.Ch.Close()

	select {}
}
