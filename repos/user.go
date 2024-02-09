package repos

import (
	"errors"
	"gorm.io/gorm"
	"main/config"
	"main/models"
	"main/utils"
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

func (r *UserRepository) Login(email, password string) (string, error) {
	if err := r.db.
		Where("email = ? AND password = ? AND otp is NULL", email, password).
		First(&models.User{}).Error; err != nil {
		return "", errors.New("Invalid credentials")
	}

	token, err := utils.GenerateToken(email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user *models.User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(user *models.User) (*models.User, error) {
	err := r.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
