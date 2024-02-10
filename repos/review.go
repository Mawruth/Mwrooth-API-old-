package repos

import (
	"main/config"
	"main/models"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository() *ReviewRepository {
	db := config.GetDB()
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) Create(Review *models.Review) (*models.Review, error) {
	err := r.db.Create(Review).Error
	if err != nil {
		return nil, err
	}
	return Review, nil
}

func (r *ReviewRepository) GetReviewByMuseum(museumId int) (*models.Review, error) {
	var Review models.Review
	err := r.db.Where("museum_id = ?", museumId).First(&Review).Error
	if err != nil {
		return nil, err
	}
	return &Review, nil
}

func (r *ReviewRepository) GetAllReviews() (*[]models.Review, error) {
	var Reviews []models.Review
	err := r.db.Find(&Reviews).Error
	if err != nil {
		return nil, err
	}
	return &Reviews, nil
}

func (r *ReviewRepository) Update(Review *models.Review) (*models.Review, error) {
	err := r.db.Save(Review).Error
	if err != nil {
		return nil, err
	}
	return Review, nil
}
