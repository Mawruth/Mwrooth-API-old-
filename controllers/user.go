package controllers

import (
	"errors"
	"main/errorHandling"
	"main/models"
	"main/services"

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
