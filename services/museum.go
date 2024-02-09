package services

import (
	"main/data/req"
	"main/models"
	"main/repos"
	"strconv"
	"strings"
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

func (m *MuseumService) GetByRating(ratingP string) ([]*models.Museum, error) {
	rating, err := strconv.ParseFloat(ratingP, 32)
	if err != nil {
		return nil, err
	}
	return m.museumRepository.GetByRating(float32(rating))
}

func (m *MuseumService) GetByTypes(typesP string) ([]*models.Museum, error) {
	typesS := strings.Split(typesP, ",")
	types := make([]int, 0)
	for _, t := range typesS {
		id, err := strconv.Atoi(t)
		if err != nil {
			return nil, err
		}
		types = append(types, id)
	}
	return m.museumRepository.GetByTypes(types)
}

func (m *MuseumService) GetByCity(city string) ([]*models.Museum, error) {
	return m.museumRepository.GetByCity(city)
}
