package service

import (
	"context"
	"time"

	"github.com/andy-ahmedov/task_manager_grpc_api/api"
	"github.com/andy-ahmedov/task_manager_server/internal/domain"
)

func (t *TasksService) CreateTask(ctx context.Context, req *api.CreateRequest) error {
	task := domain.Task{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		Created_at:  time.Now(),
	}

	return t.repo.Create(ctx, &task)

}

func (t *TasksService) GetTask(ctx context.Context, id int64) (domain.Task, error) {
	return t.repo.Get(ctx, id)
}

func (t *TasksService) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	return t.repo.GetAll(ctx)
}

func (t *TasksService) DeleteTask(ctx context.Context, id int64) error {
	return t.repo.Delete(ctx, id)
}

func (t *TasksService) UpdateTask(ctx context.Context, req *api.UpdateRequest) error {
	task := ConvertToDomainUpdateTask(req)

	return t.repo.Update(ctx, req.ID, task)
}
