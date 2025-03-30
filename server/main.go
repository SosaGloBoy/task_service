package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"os"
	"task_service/config"
	"task_service/internal/handlers"
	"task_service/internal/repository"
	"task_service/internal/routes"
	"task_service/internal/service"
)

func main() {

	cfg := config.LoadConfig()

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
	}))

	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Error initializing tasks database: %v", err)
	}

	taskRepo := repository.NewTaskRepository(db, logger)

	taskService := service.NewTaskService(taskRepo, logger)

	taskHandler := handlers.NewTaskHandler(taskService, logger)

	router := gin.Default()
	routes.SetupRoutes(router, taskHandler)

	log.Println("Starting task service on port 8086")
	if err := router.Run(":8086"); err != nil {
		logger.Error("Failed to start server", "error", err)
	}
}
