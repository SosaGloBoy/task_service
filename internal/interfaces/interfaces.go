package interfaces

import (
	"context"
	"task_service/internal/model"
)

type TaskInterface interface {
	CreateTask(ctx context.Context, task *model.Task) error
	GetTaskByID(ctx context.Context, id uint) (*model.Task, error)
	GetAllTasks(ctx context.Context) ([]*model.Task, error)
	UpdateTask(ctx context.Context, task *model.Task) error
	DeleteTask(ctx context.Context, id uint) error
}
