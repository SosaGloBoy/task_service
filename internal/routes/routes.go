package routes

import (
	"github.com/gin-gonic/gin"
	"task_service/internal/handlers"
)

func SetupRoutes(router *gin.Engine, taskHandler *handlers.TaskHandler) {
	taskGroup := router.Group("/tasks")
	{
		taskGroup.POST("", taskHandler.CreateTaskHandler)

		taskGroup.GET("/:id", taskHandler.GetTaskByIDHandler)

		taskGroup.GET("", taskHandler.GetAllTasksHandler)

		taskGroup.PUT("/:id", taskHandler.UpdateTaskHandler)

		taskGroup.DELETE("/:id", taskHandler.DeleteTaskHandler)
	}
}
