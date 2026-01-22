package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "secrets-manager-platform/docs"
	"secrets-manager-platform/internal/handlers"
	"secrets-manager-platform/internal/services"
)

// @title           AWS Secrets Manager Platform API
// @version         1.0
// @description     A REST API for managing AWS Secrets Manager secrets.
// @description     This API allows you to list, view, create, and update secrets stored in AWS Secrets Manager.

// @contact.name   API Support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api

// @externalDocs.description  AWS Secrets Manager Documentation
// @externalDocs.url          https://docs.aws.amazon.com/secretsmanager/
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	secretsService, err := services.NewSecretsService()
	if err != nil {
		log.Fatalf("Failed to initialize secrets service: %v", err)
	}

	handler := handlers.NewSecretsHandler(secretsService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Swagger documentation endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		api.GET("/secrets", handler.ListSecrets)
		api.GET("/secrets/*name", handler.GetSecret)
		api.POST("/secrets", handler.CreateSecret)
		api.PUT("/secrets/*name", handler.UpdateSecret)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Printf("Swagger documentation available at http://localhost:%s/swagger/index.html", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
