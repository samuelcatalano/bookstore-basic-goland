package repository

import (
	"bookstore/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) Create(book *models.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.DB.Find(&books).Error
	return books, err
}

// Implement other CRUD operations (GetByID, Update, Delete)
