package services

import (
	"main/models"
	"main/repos"
)

type ReviewService struct {
	ReviewRepository *repos.ReviewRepository
}

func NewReviewService() *ReviewService {
	return &ReviewService{
		ReviewRepository: repos.NewReviewRepository(),
	}
}

func (s *ReviewService) Create(Review *models.Review) (*models.Review, error) {
	return s.ReviewRepository.Create(Review)
}

func (s *ReviewService) GetAllReviews() (*[]models.Review, error) {
	return s.ReviewRepository.GetAllReviews()
}

func (s *ReviewService) GetReviewByMuseum(museumId int) (*models.Review, error) {
	return s.ReviewRepository.GetReviewByMuseum(museumId)
}

func (s *ReviewService) Update(Review *models.Review) (*models.Review, error) {
	return s.ReviewRepository.Update(Review)
}
