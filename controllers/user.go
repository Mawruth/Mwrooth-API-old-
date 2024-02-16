package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/h2non/filetype"
	"gorm.io/gorm"
	"io"
	"main/errorHandling"
	"main/models"
	"main/services"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	userService := services.NewUserService()
	return &UserController{
		userService: userService,
	}
}

func SetupUserRoutes(router fiber.Router) {
	userController := NewUserController()
	router.Get("/:id", userController.GetUser)
	router.Patch("/:id", userController.UpdateUser)
	router.Get("/:email", userController.GetUserByEmail)
	router.Post("/register", errorHandling.ValidateRegister, userController.Register)
	router.Post("/login", userController.Login)
	router.Post("/otp/verify", userController.VerifyOTP)
	router.Post("/otp/resend", userController.ResendOTP)
}

func (uc *UserController) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	user, err := uc.userService.GetUser(id)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(user)
}

func (uc *UserController) Register(c *fiber.Ctx) error {
	var user *models.User
	c.BodyParser(&user)
	result, err := uc.userService.CreateUser(user)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	return c.JSON(result)
}

func (uc *UserController) Login(c *fiber.Ctx) error {
	var user *models.User
	if err := c.BodyParser(&user); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	if user.Email == "" || user.Password == "" {
		return errorHandling.HandleHTTPError(c, errors.New("email and password are required"))
	}
	result, err := uc.userService.Login(user.Email, user.Password)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	return c.JSON(result)
}

func (uc *UserController) VerifyOTP(c *fiber.Ctx) error {
	var body fiber.Map
	if err := c.BodyParser(&body); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	if _, ok := body["email"].(string); !ok {
		return errorHandling.HandleHTTPError(c, errors.New("Enter a valid email"))
	}
	if _, ok := body["otp"].(string); !ok {
		return errorHandling.HandleHTTPError(c, errors.New("Enter a valid otp"))
	}
	if err := uc.userService.VerifyOTP(body["email"].(string), body["otp"].(string)); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "OTP verified",
	})
}

func (uc *UserController) ResendOTP(c *fiber.Ctx) error {
	var body fiber.Map
	if err := c.BodyParser(&body); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	if err := uc.userService.ResendOTP(body["email"].(string)); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "OTP resent",
	})
}

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	user := &models.UpdateUserDto{}
	newUser := &models.User{
		Model: &gorm.Model{},
	}

	if err := c.BodyParser(user); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	avatar, err := c.FormFile("avatar")
	if err == nil {
		avatarFile, err := avatar.Open()
		if err != nil {
			return errorHandling.HandleHTTPError(c, err)
		}
		defer avatarFile.Close()
		avatarUrl, err := uploadImageToS3(avatarFile)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "Failed to upload image",
			})
		}
		newUser.Avatar = avatarUrl
	}

	newUser.ID = uint(id)
	newUser.FullName = user.FullName
	newUser.Email = user.Email
	newUser.Password = user.Password
	result, err := uc.userService.UpdateUser(newUser)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	return c.JSON(result)
}

func (uc *UserController) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := uc.userService.GetUserByEmail(email)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(user)
}

func uploadImageToS3(file io.Reader) (string, error) {
	allowedExtensions := []string{"jpg", "jpeg", "png"}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("S3_REGION")),
	})
	if err != nil {
		return "", err
	}
	fmt.Println(sess.Config.Credentials.Get())

	buf, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	if len(buf) > 1e7 {
		return "", fmt.Errorf("File size is exceeds 2MB limit")
	}

	kind, _ := filetype.Match(buf)
	if !contains(allowedExtensions, kind.Extension) {
		return "", fmt.Errorf("Unsupported file type. please upload jpg, jpeg or png")
	}

	uploader := s3manager.NewUploader(sess)
	extension := kind.Extension
	fileName := generateUniqueTimestamp() + "." + extension

	upParams := &s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("S3_BUCKET")),
		Key:         aws.String(fmt.Sprintf("%s", fileName)),
		Body:        bytes.NewReader(buf),
		ContentType: aws.String(kind.MIME.Value),
		ACL:         aws.String("public-read"),
	}

	result, err := uploader.Upload(upParams)
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	return result.Location, nil
}

func generateUniqueTimestamp() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
