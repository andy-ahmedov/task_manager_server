package mongodb

import (
	"context"

	"github.com/andy-ahmedov/task_manager_server/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogItemRepository struct {
	db *mongo.Database
}

func NewLogItemRepository(db *mongo.Database) *LogItemRepository {
	return &LogItemRepository{
		db: db,
	}
}

func (l *LogItemRepository) Insert(ctx context.Context, item domain.LogItem) error {
	_, err := l.db.Collection("logs").InsertOne(ctx, item)

	return err
}

func (l *LogItemRepository) Get(ctx context.Context, id int) (*domain.LogItem, error) {
	panic("implement me")
}
