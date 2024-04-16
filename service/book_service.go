package service

import (
	"bookstore/models"
	"bookstore/repository"
)

type BookService struct {
	Repository repository.BookRepository
}

func (s *BookService) Create(book *models.Book) error {
	// Perform any necessary validation or processing
	return s.Repository.Create(book)
}

func (s *BookService) GetAll() ([]models.Book, error) {
	return s.Repository.GetAll()
}

// Implement other service methods (GetByID, Update, Delete)
