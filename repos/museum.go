package repos

import (
	"main/config"
	"main/models"

	"gorm.io/gorm"
)

type MuseumRepository struct {
	db *gorm.DB
}

func NewMuseumRepository() *MuseumRepository {
	db := config.GetDB()
	return &MuseumRepository{db: db}
}

func (r *MuseumRepository) Create(museum *models.Museum) (*models.Museum, error) {
	err := r.db.Create(museum).Error
	if err != nil {
		return nil, err
	}
	return museum, nil
}

func (r *MuseumRepository) GetAll() ([]*models.Museum, error) {
	var museums []*models.Museum
	return museums, r.db.Preload("Types").Preload("Images").Find(&museums).Error
}
