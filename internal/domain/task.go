package domain

import (
	"errors"
	"time"
)

type Task struct {
	ID          int64
	Name        string
	Description string
	Status      string
	Created_at  time.Time
}

type UpdateTaskInput struct {
	Name        *string
	Description *string
	Status      *string
}

var (
	ErrTaskNotFound = errors.New("Task not fount")
)
