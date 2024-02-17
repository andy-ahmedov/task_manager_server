package mongodb

import (
	"context"

	"github.com/andy-ahmedov/task_manager_server/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	db *mongo.Database
}

func (m *MongoRepository) Insert(ctx context.Context, item domain.LogItem) error {
	_, err := m.db.Collection("logs").InsertOne(ctx, item)

	return err
}

func (m *MongoRepository) Get(ctx context.Context, id int) (*domain.LogItem, error) {
	panic("implement me")
}
