package service

import (
	"context"
	"fmt"
	"log/slog"
	"task_service/internal/interfaces"
	"task_service/internal/model"
)

type TaskService struct {
	TaskRepository interfaces.TaskInterface
	Logger         *slog.Logger
}

func NewTaskService(taskRepository interfaces.TaskInterface, logger *slog.Logger) *TaskService {
	return &TaskService{
		TaskRepository: taskRepository,
		Logger:         logger,
	}
}
func (s *TaskService) CreateTask(ctx context.Context, task *model.Task) error {
	if err := s.TaskRepository.CreateTask(ctx, task); err != nil {
		return fmt.Errorf("failed to create task: %v", err)
	}
	return nil
}

// GetTaskByID возвращает задание по ID
func (s *TaskService) GetTaskByID(ctx context.Context, id uint) (*model.Task, error) {
	task, err := s.TaskRepository.GetTaskByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task by id: %v", err)
	}
	return task, nil
}

// GetAllTasks возвращает все задания
func (s *TaskService) GetAllTasks(ctx context.Context) ([]*model.Task, error) {
	tasks, err := s.TaskRepository.GetAllTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %v", err)
	}
	return tasks, nil
}

// UpdateTask обновляет задание
func (s *TaskService) UpdateTask(ctx context.Context, task *model.Task) error {
	if err := s.TaskRepository.UpdateTask(ctx, task); err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}
	return nil
}

// DeleteTask удаляет задание
func (s *TaskService) DeleteTask(ctx context.Context, id uint) error {
	if err := s.TaskRepository.DeleteTask(ctx, id); err != nil {
		return fmt.Errorf("failed to delete task: %v", err)
	}
	return nil
}
