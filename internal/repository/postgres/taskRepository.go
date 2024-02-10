package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/andy-ahmedov/task_manager_server/internal/domain"
	"github.com/jackc/pgx/v5"
)

type Task struct {
	db *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) *Task {
	return &Task{
		db: db,
	}
}

func (t *Task) Create(ctx context.Context, task *domain.Task) error {
	request := `INSERT INTO tasks(name, description, status, created_at) VALUES($1, $2, $3, $4) RETURNING id`

	err := t.db.QueryRow(ctx, request, task.Name, task.Description, task.Status, task.Created_at).Scan(&task.ID)

	return err
}

func (t *Task) Get(ctx context.Context, id int64) (domain.Task, error) {
	var task domain.Task

	request := `SELECT * FROM tasks WHERE id=$1`

	err := t.db.QueryRow(ctx, request, id).Scan(&task.ID, &task.Name, &task.Description, &task.Status, &task.Created_at)
	if err == sql.ErrNoRows {
		return task, domain.ErrTaskNotFound
	}

	return task, err
}

func (t *Task) GetAll(ctx context.Context) ([]domain.Task, error) {
	tasks := make([]domain.Task, 0)

	request := `SELECT * FROM tasks`

	rows, err := t.db.Query(ctx, request)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task domain.Task

		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Status, &task.Created_at)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *Task) Delete(ctx context.Context, id int64) error {
	request := `DELETE FROM tasks WHERE id=$1`

	_, err := t.db.Exec(ctx, request, id)

	return err
}

func (t *Task) Update(ctx context.Context, id int64, updTask domain.UpdateTaskInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updTask.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *updTask.Name)
		argId++
	}

	if updTask.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *updTask.Description)
		argId++
	}

	if updTask.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *updTask.Status)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE tasks SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, id)

	_, err := t.db.Exec(ctx, query, args...)

	return err
}
