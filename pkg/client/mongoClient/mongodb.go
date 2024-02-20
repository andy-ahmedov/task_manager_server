package mongoClient

import (
	"context"
	"fmt"

	"github.com/andy-ahmedov/task_manager_server/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, cfg config.Mongo) (db *mongo.Database, err error) {
	mongoDBurl, isAuth := CreateMongoDBURI(cfg)

	clientOptions := NewClientOptions(cfg, mongoDBurl, isAuth)

	client, err := ConnectAndPing(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return client.Database(cfg.Database), err
}

func CreateMongoDBURI(cfg config.Mongo) (string, bool) {
	var mongoDBurl string
	var isAuth bool

	if cfg.Username == "" && cfg.Password == "" {
		mongoDBurl = fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port)
	} else {
		isAuth = true
		mongoDBurl = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
		fmt.Println(mongoDBurl)
	}

	return mongoDBurl, isAuth
}

func NewClientOptions(cfg config.Mongo, mongoDBurl string, isAuth bool) *options.ClientOptions {
	clientOptions := options.Client().ApplyURI(mongoDBurl)
	// if isAuth {
	// 	if cfg.AuthDB == "" {
	// 		cfg.AuthDB = cfg.Database
	// 		fmt.Println(cfg.AuthDB)
	// 	}

	// 	clientOptions.SetAuth(options.Credential{
	// 		AuthSource: cfg.AuthDB,
	// 		Username:   cfg.Username,
	// 		Password:   cfg.Password,
	// 	})
	// }

	return clientOptions
}

func ConnectAndPing(ctx context.Context, clientOptions *options.ClientOptions) (*mongo.Client, error) {
	// client, err := mongo.Connect(ctx, clientOptions)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		// return nil, fmt.Errorf("failed to connect to mongoDB due to error: %v", err)
		return nil, fmt.Errorf("Error while try to initialize mongo db %v", err)
	}

	// Create connect
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error while connect to mongodb %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB due to error: %v", err)
	}

	return client, err
}
