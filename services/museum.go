package services

import (
	"main/data/req"
	"main/models"
	"main/repos"
)

type MuseumService struct {
	museumRepository *repos.MuseumRepository
}

func NewMuseumService() *MuseumService {
	return &MuseumService{
		museumRepository: repos.NewMuseumRepository(),
	}
}

func (m *MuseumService) CreateMuseum(museum req.MuseumReq) (*models.Museum, error) {

	// find all types by ids from database
	var types []models.Type
	for _, id := range museum.Types {
		type_, _ := NewTypeService().typeRepository.GetById(id)
		types = append(types, type_)
	}

	// add all images
	var images []models.MuseumImages
	for _, image := range museum.Images {
		images = append(images, models.MuseumImages{
			Image_path: image,
		})
	}

	newMuseum := models.Museum{
		Name:        museum.Name,
		Description: museum.Description,
		WorkTime:    museum.WorkTime,
		Country:     museum.Country,
		City:        museum.City,
		Street:      museum.Street,
		Types:       types,
		Images:      images,
	}

	return m.museumRepository.Create(&newMuseum)
}

func (m *MuseumService) GetAll() ([]*models.Museum, error) {
	return m.museumRepository.GetAll()
}
