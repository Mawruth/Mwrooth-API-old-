package repos

import (
	"main/config"
	"main/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	db := config.GetDB()
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *models.Category) (*models.Category, error) {
	err := r.db.Create(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}