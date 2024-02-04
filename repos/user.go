package repos

import (
	"gorm.io/gorm"
	"main/config"
	"main/models"
	"sync"
)

var (
	once sync.Once
	repo *UserRepository
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	db := config.GetDB()
	once.Do(func() {
		repo = &UserRepository{
			db: db,
		}
	})

	return repo
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	var user *models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
