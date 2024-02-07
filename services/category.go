package services

import (
	"main/models"
	"main/repos"
)

type CategoryService struct {
	categoryRepository *repos.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		categoryRepository: repos.NewCategoryRepository(),
	}
}

func (s *CategoryService) Create(category *models.Category) (*models.Category, error) {
	return s.categoryRepository.Create(category)
}