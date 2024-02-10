package repos

import (
	"main/config"
	"main/models"

	"gorm.io/gorm"
)

type StoryRepository struct {
	db *gorm.DB
}

func NewStoryRepository() *StoryRepository {
	db := config.GetDB()
	return &StoryRepository{db: db}
}

func (r *StoryRepository) Create(Story *models.Story) (*models.Story, error) {
	err := r.db.Create(Story).Error
	if err != nil {
		return nil, err
	}
	return Story, nil
}

func (r *StoryRepository) GetAllStories() (*[]models.Story, error) {
	var Stories []models.Story
	err := r.db.Find(&Stories).Error
	if err != nil {
		return nil, err
	}
	return &Stories, nil
}