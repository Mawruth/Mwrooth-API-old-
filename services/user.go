package services

import (
	"errors"
	"main/data/res"
	"main/models"
	"main/repos"
	"main/utils"
	"os"
	"strconv"
	"sync"
	"time"

	"gopkg.in/gomail.v2"
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
	rawOTP := utils.GenerateOTP()
	otp := rawOTP + "|1|"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	otp += timestamp
	user.OTP = &otp
	verificationMsg := utils.GenerateVerificationMessage(rawOTP)
	if err := u.SendCode(user.Email, verificationMsg); err != nil {
		return nil, err
	}
	return u.userRepository.Create(user)
}

func (u *UserService) Login(email, password string) (*res.UserRes, error) {
	return u.userRepository.Login(email, password)
}

func (u *UserService) SendCode(email string, verificationMsg []byte) error {
	host := os.Getenv("SMTP_HOST")
	port, _ := strconv.ParseInt(os.Getenv("SMTP_PORT"), 10, 32)
	senderEmail := os.Getenv("SENDER_EMAIL")
	mail := gomail.NewMessage()
	mail.SetHeader("From", senderEmail)
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "Verification Code")
	mail.SetBody("text/html", string(verificationMsg))
	dialer := gomail.NewDialer(host, int(port), senderEmail, os.Getenv("SENDER_PASSWORD"))
	if err := dialer.DialAndSend(mail); err != nil {
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

	rawOTP := utils.GenerateOTP()
	otp := rawOTP + "|1|"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	otp += timestamp
	user.OTP = &otp
	verificationMsg := utils.GenerateVerificationMessage(rawOTP)
	if err := u.SendCode(user.Email, verificationMsg); err != nil {
		return err
	}

	if _, err := u.userRepository.Update(user); err != nil {
		return err
	}

	return nil
}
