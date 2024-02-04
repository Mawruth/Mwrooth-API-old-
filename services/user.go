package services

import (
	"main/models"
	"main/repos"
	"sync"
)

var (
	once    sync.Once
	service *UserService
)

type UserService struct {
	userRepository *repos.UserRepository
}

func NewUserService() *UserService {
	userRepository := repos.NewUserRepository()
	once.Do(func() {
		service = &UserService{userRepository: userRepository}
	})

	return service
}

func (u *UserService) GetUser(id int) (*models.User, error) {
	return u.userRepository.GetByID(id)
}

func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	return u.userRepository.Create(user)
}
