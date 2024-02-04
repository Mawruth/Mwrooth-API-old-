package repos

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"main/config"
	"main/models"
	"os"
	"sync"
	"time"
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
		Where("email = ? AND password = ?", email, password).
		First(&models.User{}).Error; err != nil {
		return "", err
	}

	token, err := generateToken(email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  "USER",
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	},
	)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
