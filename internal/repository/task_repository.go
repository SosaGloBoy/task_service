package repository

import (
	"context"
	"gorm.io/gorm"
	"log/slog"
	"task_service/internal/interfaces"
	"task_service/internal/model"
)

type TaskRepository struct {
	DB     *gorm.DB
	Logger *slog.Logger
}

func NewTaskRepository(db *gorm.DB, logger *slog.Logger) interfaces.TaskInterface {
	return &TaskRepository{DB: db, Logger: logger}
}

// CreateTask создаёт задание в базе данных
func (r *TaskRepository) CreateTask(ctx context.Context, task *model.Task) error {
	r.Logger.InfoContext(ctx, "Creating task", "task", task.Title)

	if err := r.DB.Create(task).Error; err != nil {
		r.Logger.ErrorContext(ctx, "Failed to create task", "error", err)
		return err
	}
	r.Logger.InfoContext(ctx, "Task created successfully", "task", task.Title)
	return nil
}

// GetTaskByID находит задание по ID
func (r *TaskRepository) GetTaskByID(ctx context.Context, id uint) (*model.Task, error) {
	r.Logger.InfoContext(ctx, "Fetching task by ID", "id", id)
	var task model.Task
	if err := r.DB.Preload("Steps").First(&task, id).Error; err != nil {
		r.Logger.ErrorContext(ctx, "Failed to find task by ID", "id", id, "error", err)
		return nil, err
	}
	r.Logger.InfoContext(ctx, "Task found", "id", id)
	return &task, nil
}

// GetAllTasks находит все задания
func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]*model.Task, error) {
	r.Logger.InfoContext(ctx, "Fetching all tasks")
	var tasks []*model.Task
	if err := r.DB.Find(&tasks).Error; err != nil {
		r.Logger.ErrorContext(ctx, "Failed to fetch tasks", "error", err)
		return nil, err
	}
	r.Logger.InfoContext(ctx, "Tasks fetched successfully", "count", len(tasks))
	return tasks, nil
}

// UpdateTask обновляет задание
func (r *TaskRepository) UpdateTask(ctx context.Context, task *model.Task) error {
	r.Logger.InfoContext(ctx, "Updating task", "task", task.Title)
	if err := r.DB.Save(task).Error; err != nil {
		r.Logger.ErrorContext(ctx, "Failed to update task", "task", task.Title, "error", err)
		return err
	}
	r.Logger.InfoContext(ctx, "Task updated successfully", "task", task.Title)
	return nil
}

// DeleteTask удаляет задание по ID
func (r *TaskRepository) DeleteTask(ctx context.Context, id uint) error {
	r.Logger.InfoContext(ctx, "Deleting task", "id", id)
	if err := r.DB.Delete(&model.Task{}, id).Error; err != nil {
		r.Logger.ErrorContext(ctx, "Failed to delete task", "id", id, "error", err)
		return err
	}
	r.Logger.InfoContext(ctx, "Task deleted successfully", "id", id)
	return nil
}
