package main

import (
	"context"
	"log"

	"github.com/andy-ahmedov/task_manager_server/internal/config"
	"github.com/andy-ahmedov/task_manager_server/internal/logger"
	"github.com/andy-ahmedov/task_manager_server/internal/repository/mongodb"
	"github.com/andy-ahmedov/task_manager_server/internal/repository/postgres"
	"github.com/andy-ahmedov/task_manager_server/internal/service/logItemService"
	"github.com/andy-ahmedov/task_manager_server/internal/service/taskService"
	grpc_client "github.com/andy-ahmedov/task_manager_server/internal/transport/grpc"
	"github.com/andy-ahmedov/task_manager_server/internal/transport/rabbitmq"
	"github.com/andy-ahmedov/task_manager_server/pkg/client/mongoClient"
	"github.com/andy-ahmedov/task_manager_server/pkg/client/postgresClient"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	logg := logger.NewLogger()
	logg.Info(cfg)

	postgresDB, err := postgresClient.ConnectToDB(cfg.PostgresDB)
	if err != nil {
		logg.Fatal(err)
	}

	mongoClient, err := mongoClient.NewClient(context.Background(), cfg.MongoDB)
	if err != nil {
		logg.Fatal(err)
	}

	postgresRepository := postgres.NewTaskRepository(postgresDB)
	mongoRepository := mongodb.NewLogItemRepository(mongoClient)

	logItemService := logItemService.NewLogItemsService(mongoRepository)
	taskService := taskService.NewTaskService(postgresRepository)

	taskSrv := grpc_client.NewCreaterServer(taskService, logg)
	srv := grpc_client.New(taskSrv)

	broker := rabbitmq.NewBroker(&cfg.Brkr, logg, logItemService)

	go broker.RunBrokerServer(context.Background())

	if err := srv.ListenAndServe(cfg.Srvr.Port); err != nil {
		logg.Fatal(err)
	}
}
