package main

import (
	"log"
	"os"

	"user-service/controllers"
	"user-service/models"
	"user-service/repository"
	"user-service/services"

	"github.com/gin-gonic/gin"
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

	// Migrate the schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	// Setup clean architecture components
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo, os.Getenv("JWT_SECRET"))
	userController := controllers.NewUserController(userService)

	// Setup Gin router
	r := gin.Default()

	// Routes
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/register", userController.Register)
		authRoutes.POST("/login", userController.Login)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
