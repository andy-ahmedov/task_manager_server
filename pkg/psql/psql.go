package psql

import (
	"context"
	"fmt"
	"time"

	"github.com/andy-ahmedov/task_manager_server/internal/config"
	"github.com/jackc/pgx/v5"
)

func ConnectToDB(cfg config.Postgres) (*pgx.Conn, error) {
	conn_str := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	var conn *pgx.Conn
	var err error

	fmt.Println(conn_str)

	for i := 0; i < 10; i++ {
		conn, err := pgx.Connect(context.Background(), conn_str)
		if err != nil {
			fmt.Println("TRYING №", i, err)
			time.Sleep(time.Second)
			continue
			// return nil, err
		}

		err = conn.Ping(context.TODO())
		if err != nil {
			fmt.Println("TRYING №", i, err)
			time.Sleep(time.Second)
			continue
			// return nil, err
		}
		break
	}
	if err != nil {
		fmt.Println("Failed to connect to database after 10 attempts")
		return nil, err
	} else {
		fmt.Println("CONNECTED")
		return conn, nil
	}
}
