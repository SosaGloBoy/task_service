package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
	"task_service/internal/model"
	"task_service/internal/service"
)

type TaskHandler struct {
	TaskService *service.TaskService
	Logger      *slog.Logger
}

// NewTaskHandler создаёт новый экземпляр TaskHandler
func NewTaskHandler(taskService *service.TaskService, logger *slog.Logger) *TaskHandler {
	return &TaskHandler{
		TaskService: taskService,
		Logger:      logger,
	}
}

// CreateTaskHandler создаёт задание
func (h *TaskHandler) CreateTaskHandler(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := h.TaskService.CreateTask(c, &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully", "task": task})
}

// GetTaskByIDHandler возвращает задание по ID
func (h *TaskHandler) GetTaskByIDHandler(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := h.TaskService.GetTaskByID(c, uint(taskID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

// GetAllTasksHandler возвращает все задания
func (h *TaskHandler) GetAllTasksHandler(c *gin.Context) {
	tasks, err := h.TaskService.GetAllTasks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// UpdateTaskHandler обновляет задание
func (h *TaskHandler) UpdateTaskHandler(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	task.Id = uint(taskID)

	if err := h.TaskService.UpdateTask(c, &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

// DeleteTaskHandler удаляет задание
func (h *TaskHandler) DeleteTaskHandler(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := h.TaskService.DeleteTask(c, uint(taskID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
