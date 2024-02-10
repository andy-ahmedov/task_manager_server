package psql

import (
	"context"
	"fmt"

	"github.com/andy-ahmedov/task_manager_server/internal/config"
	"github.com/jackc/pgx/v5"
)

func ConnectToDB(cfg config.Postgres) (*pgx.Conn, error) {
	conn_str := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	fmt.Println(conn_str)
	conn, err := pgx.Connect(context.Background(), conn_str)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.TODO())
	if err != nil {
		return nil, err
	}

	fmt.Println("CONNECTED")

	return conn, nil
}
