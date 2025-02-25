package repos

import (
	"errors"
	"main/config"
	"main/data/res"
	"main/models"
	"main/utils"
	"sync"

	"gorm.io/gorm"
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

func (r *UserRepository) Login(email, password string) (*res.UserRes, error) {
	user := &models.User{}
	if err := r.db.
		Where("email = ? AND password = ? AND otp is NULL", email, password).
		First(user).Error; err != nil {
		return nil, errors.New("Invalid credentials")
	}

	token, err := utils.GenerateToken(email)
	if err != nil {
		return nil, err
	}

	userRes := res.UserRes{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		UserName: user.UserName,
		Token:    token,
	}

	return &userRes, nil
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
	res := r.db.Save(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}
