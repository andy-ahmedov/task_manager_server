package logItemService

import (
	"context"

	"github.com/andy-ahmedov/task_manager_server/internal/domain"
)

func (l *LogItemsService) Create(ctx context.Context, logItem domain.LogItem) error {
	err := l.repo.Insert(ctx, logItem)

	return err
}
