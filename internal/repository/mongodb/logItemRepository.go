package mongodb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/andy-ahmedov/task_manager_server/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	nextID, err := l.getNextSequence(ctx, "logitem_id")
	if err != nil {
		return fmt.Errorf("failed to get next id: %v", err)
	}

	item.ID = nextID
	_, err = l.db.Collection("logs").InsertOne(ctx, item)
	if err != nil {
		return fmt.Errorf("failed to create user due to error: %v", err)
	}

	return err
}

func (l *LogItemRepository) getNextSequence(ctx context.Context, sequenceName string) (string, error) {
	result := l.db.Collection("counters").FindOneAndUpdate(
		ctx,
		bson.M{"_id": sequenceName},
		bson.M{"$inc": bson.M{"seq": 1}},
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	)

	var counter struct {
		Seq int
	}
	if err := result.Decode(&counter); err != nil {
		return "", fmt.Errorf("failed to decode counter: %v", err)
	}

	return strconv.Itoa(counter.Seq), nil
}

func (l *LogItemRepository) Get(ctx context.Context, id int) (*domain.LogItem, error) {
	panic("implement me")
}
