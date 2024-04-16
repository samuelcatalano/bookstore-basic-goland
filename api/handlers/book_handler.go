package handlers

import (
	"bookstore/models"
	"bookstore/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookHandler struct {
	Service service.BookService
}

func (h *BookHandler) Create(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.Create(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) GetAll(c *gin.Context) {
	books, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// Implement other handler methods (GetByID, Update, Delete)
