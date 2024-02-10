package service

import (
	"context"

	"github.com/andy-ahmedov/task_manager_server/internal/domain"
	"github.com/andy-ahmedov/task_manager_grpc_api/api"
)

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task) error
	Get(ctx context.Context, id int64) (domain.Task, error)
	GetAll(ctx context.Context) ([]domain.Task, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, updTask domain.UpdateTaskInput) error
}

type TasksStorage struct {
	repo TaskRepository
}

func NewTaskStorage(repo TaskRepository) *TasksStorage {
	return &TasksStorage{
		repo: repo,
	}
}

func ConvertToDomainUpdateTask(req *api.UpdateRequest) domain.UpdateTaskInput {
	task := domain.UpdateTaskInput{}

	if req.Task.Name != nil {
		task.Name = &req.Task.Name.Value
	} else {
		task.Name = nil
	}

	if req.Task.Description != nil {
		task.Description = &req.Task.Description.Value
	} else {
		task.Description = nil
	}

	if req.Task.Status != nil {
		task.Status = &req.Task.Status.Value
	} else {
		task.Status = nil
	}

	return task
}
