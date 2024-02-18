package logItemService

import (
	"context"

	"github.com/andy-ahmedov/task_manager_server/internal/domain"
)

type LogItemRepository interface {
	Insert(ctx context.Context, item domain.LogItem) error
}

type LogItemsService struct {
	repo LogItemRepository
}

func NewLogItemsService(repo LogItemRepository) *LogItemsService {
	return &LogItemsService{
		repo: repo,
	}
}
