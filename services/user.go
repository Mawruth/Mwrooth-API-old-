package services

import (
	"errors"
	"main/models"
	"main/repos"
	"main/utils"
	"net"
	"net/smtp"
	"os"
	"sync"
	"time"
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
	otp := utils.GenerateOTP() + "|1|"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	otp += timestamp
	user.OTP = &otp
	//verificationMsg := utils.GenerateVerificationMessage(user.OTP)
	//if err := u.SendCode(user.Email, verificationMsg); err != nil {
	//	return nil, err
	//}
	return u.userRepository.Create(user)
}

func (u *UserService) Login(email, password string) (string, error) {
	return u.userRepository.Login(email, password)
}

func (u *UserService) SendCode(email string, verificationMsg []byte) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	senderEmail := os.Getenv("SENDER_EMAIL")
	auth := smtp.PlainAuth("", senderEmail, os.Getenv("SENDER_PASSWORD"), host)
	if err := smtp.SendMail(
		net.JoinHostPort(host, port),
		auth,
		senderEmail,
		[]string{email},
		verificationMsg,
	); err != nil {
		return err
	}

	return nil
}

func (u *UserService) VerifyOTP(email, otp string) error {
	user, err := u.userRepository.GetByEmail(email)
	if err != nil {
		return err
	}

	rawOTP, timestampAsTime := utils.ParseOTP(*user.OTP)

	if time.Now().After(timestampAsTime) {
		return errors.New("OTP expired")
	}

	if rawOTP != otp {
		return errors.New("Invalid otp")
	}

	user.OTP = nil
	if _, err := u.userRepository.Update(user); err != nil {
		return err
	}

	return nil
}

func (u *UserService) ResendOTP(email string) error {
	user, err := u.userRepository.GetByEmail(email)
	if err != nil {
		return err
	}

	otp := utils.GenerateOTP() + "|1|"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	otp += timestamp
	user.OTP = &otp
	//verificationMsg := utils.GenerateVerificationMessage(user.OTP)
	//if err := u.SendCode(user.Email, verificationMsg); err != nil {
	//	return nil, err
	//}

	if _, err := u.userRepository.Update(user); err != nil {
		return err
	}

	return nil
}
