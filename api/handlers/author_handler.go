package handlers

import (
	"bookstore/models"
	"bookstore/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthorHandler struct {
	Service service.AuthorService
}

func (h *AuthorHandler) Create(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.Create(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, author)
}

func (h *AuthorHandler) GetAll(c *gin.Context) {
	authors, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authors)
}

// Implement other handler methods (GetByID, Update, Delete)
