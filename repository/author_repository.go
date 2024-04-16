package repository

import (
	"bookstore/models"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	DB *gorm.DB
}

func (r *AuthorRepository) Create(author *models.Author) error {
	return r.DB.Create(author).Error
}

func (r *AuthorRepository) GetAll() ([]models.Author, error) {
	var authors []models.Author
	err := r.DB.Find(&authors).Error
	return authors, err
}

// Implement other CRUD operations (GetByID, Update, Delete)
