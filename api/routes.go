package api

import (
	"bookstore/api/handlers"
	"bookstore/repository"
	"bookstore/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, router *gin.Engine) {
	authorRepo := repository.AuthorRepository{DB: db}
	authorService := service.AuthorService{Repository: authorRepo}
	authorHandler := handlers.AuthorHandler{Service: authorService}

	bookRepo := repository.BookRepository{DB: db}
	bookService := service.BookService{Repository: bookRepo}
	bookHandler := handlers.BookHandler{Service: bookService}

	router.POST("/authors", authorHandler.Create)
	router.GET("/authors", authorHandler.GetAll)
	// Add other routes for author (GetByID, Update, Delete)

	router.POST("/books", bookHandler.Create)
	router.GET("/books", bookHandler.GetAll)
	// Add other routes for book (GetByID, Update, Delete)
}
