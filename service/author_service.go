package service

import (
	"bookstore/models"
	"bookstore/repository"
)

type AuthorService struct {
	Repository repository.AuthorRepository
}

func (s *AuthorService) Create(author *models.Author) error {
	// Perform any necessary validation or processing
	return s.Repository.Create(author)
}

func (s *AuthorService) GetAll() ([]models.Author, error) {
	return s.Repository.GetAll()
}

// Implement other service methods (GetByID, Update, Delete)
