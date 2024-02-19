package main

import (
	"log"

	"github.com/andy-ahmedov/task_manager_server/internal/config"
	"github.com/andy-ahmedov/task_manager_server/internal/logger"
	"github.com/andy-ahmedov/task_manager_server/internal/repository/postgres"
	"github.com/andy-ahmedov/task_manager_server/internal/service/taskService"
	grpc_client "github.com/andy-ahmedov/task_manager_server/internal/transport/grpc"
	"github.com/andy-ahmedov/task_manager_server/internal/transport/rabbitmq"
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

	// mongoClient, err := mongoClient.NewClient(context.Background(), cfg.MongoDB)
	// if err != nil {
	// 	logg.Fatal(err)
	// }

	postgresRepository := postgres.NewTaskRepository(postgresDB)
	// mongoRepository := mongodb.NewLogItemRepository(mongoClient)

	// _ = logItemService.NewLogItemsService(mongoRepository)
	taskService := taskService.NewTaskService(postgresRepository)

	taskSrv := grpc_client.NewCreaterServer(taskService, logg)
	srv := grpc_client.New(taskSrv)

	// if err = rabbitmq.InitRabbitMQ(&cfg.Brkr); err != nil {
	// 	logg.Fatal(err)
	// }
	go rabbitmq.InitRabbitMQ(&cfg.Brkr, logg)

	if err := srv.ListenAndServe(cfg.Srvr.Port); err != nil {
		logg.Fatal(err)
	}
}
