package main

import (
	"bookstore/api"
	"bookstore/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connect to the PostgreSQL database
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Author{}, &models.Book{})

	// Initialize the Gin router
	router := gin.Default()

	// Set up the routes
	api.SetupRoutes(db, router)

	// Start the server
	router.Run(":8080")
}
