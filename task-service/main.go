package main

import (
	"log"
	"os"

	"task-service/controllers"
	"task-service/middlewares"
	"task-service/models"
	"task-service/repository"
	"task-service/services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DB Connection string
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	// Redis connection
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})

	// Clean architecture
	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo, redisClient)
	taskController := controllers.NewTaskController(taskService)

	// Gin Router
	r := gin.Default()

	api := r.Group("/api/tasks")
	api.Use(middlewares.AuthMiddleware(redisClient))
	{
		api.POST("", taskController.CreateTask)
		api.GET("", taskController.GetTasks)
		api.PUT("/:id", taskController.UpdateTask)
		api.DELETE("/:id", taskController.DeleteTask)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
