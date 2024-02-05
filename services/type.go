package services

import (
	"main/models"
	"main/repos"
)

type TypeService struct {
	typeRepository *repos.TypeRepository
}

func NewTypeService() *TypeService {
	typeRepository := repos.NewTypeRepository()
	return &TypeService{
		typeRepository: typeRepository,
	}
}

func (t *TypeService) Create(type_ *models.Type) (*models.Type, error) {
	return t.typeRepository.Create(type_)
}

func (t *TypeService) GetAllTypes() ([]*models.Type, error) {
	return t.typeRepository.GetAll()
}

//get type by id
func (t *TypeService) GetTypeById(id int) (models.Type, error) {
	return t.typeRepository.GetById(id)
}