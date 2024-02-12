package main

import (
	"log"

	"github.com/andy-ahmedov/task_manager_server/internal/config"
	"github.com/andy-ahmedov/task_manager_server/internal/logger"
	"github.com/andy-ahmedov/task_manager_server/internal/repository/postgres"
	"github.com/andy-ahmedov/task_manager_server/internal/service"
	grpc_client "github.com/andy-ahmedov/task_manager_server/internal/transport/grpc"
	"github.com/andy-ahmedov/task_manager_server/pkg/psql"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	logg := logger.NewLogger()
	logg.Info(cfg)

	db, err := psql.ConnectToDB(cfg.DB)
	if err != nil {
		logg.Fatal(err)
	}

	postgres := postgres.NewTaskRepository(db)
	service := service.NewTaskStorage(postgres)
	taskSrv := grpc_client.NewCreaterServer(service, logg)
	srv := grpc_client.New(taskSrv)

	if err := srv.ListenAndServe(cfg.Srvr.Port); err != nil {
		logg.Fatal(err)
	}
}
