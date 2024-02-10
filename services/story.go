package services

import (
	"main/models"
	"main/repos"
)

type StoryService struct {
	StoryRepository *repos.StoryRepository
}

func NewStoryService() *StoryService {
	return &StoryService{
		StoryRepository: repos.NewStoryRepository(),
	}
}

func (s *StoryService) Create(Story *models.Story) (*models.Story, error) {
	return s.StoryRepository.Create(Story)
}

func (s *StoryService) GetAllStories() (*[]models.Story, error) {
	return s.StoryRepository.GetAllStories()
}
