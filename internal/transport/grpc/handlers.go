package grpc

import (
	"context"

	"github.com/andy-ahmedov/task_manager_grpc_api/api"
)

func (h *TaskServiceServer) Create(ctx context.Context, req *api.CreateRequest) (*api.Empty, error) {
	err := h.service.CreateTask(ctx, req)

	return &api.Empty{}, err
}

func (h *TaskServiceServer) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	task, err := h.service.GetTask(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	apiTask := ConvertToApiTask(task)

	resp := api.GetResponse{
		Task: apiTask,
	}

	return &resp, err
}

func (h *TaskServiceServer) GetAll(ctx context.Context, req *api.GetAllRequest) (*api.GetAllResponse, error) {
	tasks, err := h.service.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	apiTasks := make([]*api.Task, 0)

	for _, task := range tasks {
		apiTask := ConvertToApiTask(task)
		apiTasks = append(apiTasks, apiTask)
	}

	res := api.GetAllResponse{
		Tasks: apiTasks,
	}

	return &res, nil
}

func (h *TaskServiceServer) Delete(ctx context.Context, req *api.DeleteRequest) (*api.Empty, error) {
	err := h.service.DeleteTask(ctx, req.ID)

	return &api.Empty{}, err
}

func (h *TaskServiceServer) Update(ctx context.Context, req *api.UpdateRequest) (*api.Empty, error) {
	err := h.service.UpdateTask(ctx, req)

	return &api.Empty{}, err
}
