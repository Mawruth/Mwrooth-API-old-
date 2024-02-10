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

func (r *MuseumRepository) GetByRating(rating float32) ([]*models.Museum, error) {
	var museums []*models.Museum
	return museums, r.db.Where("rating = ?", rating).Find(&museums).Error
}

func (r *MuseumRepository) GetByTypes(types []int) ([]*models.Museum, error) {
	var types_ []*models.Type
	err := r.db.Where("id IN ?", types).Preload("Museums").Find(&types_).Error
	if err != nil {
		return nil, err
	}
	var museums []*models.Museum
	for _, t := range types_ {
		museums = append(museums, t.Museums...)
	}
	return museums, nil
}

func (r *MuseumRepository) GetByCity(city string) ([]*models.Museum, error) {
	var museums []*models.Museum
	return museums, r.db.Where("city = ?", city).Find(&museums).Error
}
func (r *MuseumRepository) GetByID(id int) (*models.Museum, error) {
	var museum *models.Museum
	return museum, r.db.Where("id = ?", id).Preload("Types").Preload("Images").Preload("Pieces.Images").Preload("Reviews").Find(&museum).Error
}
