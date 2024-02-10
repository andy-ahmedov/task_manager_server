package grpc

import (
	"context"

	"github.com/andy-ahmedov/task_manager_server/internal/domain"
	"github.com/andy-ahmedov/task_manager_grpc_api/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskService interface {
	CreateTask(ctx context.Context, req *api.CreateRequest) error
	GetTask(ctx context.Context, id int64) (domain.Task, error)
	GetAllTasks(ctx context.Context) ([]domain.Task, error)
	DeleteTask(ctx context.Context, id int64) error
	UpdateTask(ctx context.Context, req *api.UpdateRequest) error
}

type TaskServiceServer struct {
	service TaskService
	api.UnimplementedTaskServiceServer
	logger *logrus.Logger
}

func NewCreaterServer(service TaskService, logger *logrus.Logger) *TaskServiceServer {
	return &TaskServiceServer{
		service: service,
		logger:  logger,
	}
}

func ConvertToApiTask(task domain.Task) *api.Task {
	created_at := timestamppb.New(task.Created_at)

	apitask := api.Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   created_at,
	}

	return &apitask
}
