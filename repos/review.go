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
	var creator *models.User
	err = r.db.Where("id = ?", Review.UserId).First(&creator).Error
	if err != nil {
		return nil, err
	}
	Review.Creator = creator
	return Review, nil
}

func (r *ReviewRepository) GetReviewByMuseum(museumId int) (*[]models.Review, error) {
	var Review []models.Review
	err := r.db.Preload("Creator").Where("museum_id = ?", museumId).Find(&Review).Error
	if err != nil {
		return nil, err
	}
	return &Review, nil
}

func (r *ReviewRepository) GetAllReviews() (*[]models.Review, error) {
	var Reviews []models.Review
	err := r.db.Preload("Creator").Find(&Reviews).Error
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
	var creator *models.User
	err = r.db.Where("id = ?", Review.UserId).First(&creator).Error
	if err != nil {
		return nil, err
	}
	Review.Creator = creator
	return Review, nil
}
