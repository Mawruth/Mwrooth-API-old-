package repos

import (
	"main/config"
	"main/models"

	"gorm.io/gorm"
)

type TypeRepository struct {
	db *gorm.DB
}

func NewTypeRepository() *TypeRepository {
	db := config.GetDB()
	return &TypeRepository{db: db}
}

func (r *TypeRepository) Create(type_ *models.Type) (*models.Type, error) {
	err := r.db.Create(type_).Error
	if err != nil {
		return nil, err
	}
	return type_, nil
}

func (r *TypeRepository) GetAll() ([]*models.Type, error) {
	var types []*models.Type
	err := r.db.Find(&types).Error
	if err != nil {
		return nil, err
	}
	return types, nil
}

func (r *TypeRepository) GetById(id int) (models.Type, error) {
	var type_ models.Type
	err := r.db.Where("id = ?", id).First(&type_).Error
	if err != nil {
		return models.Type{}, err
	}
	return type_, nil
}
