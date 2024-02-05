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
	return c.JSON(map[string]string{
		"token": result,
	})
}
