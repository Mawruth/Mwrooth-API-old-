package controllers

import (
	"github.com/gofiber/fiber/v2"
	"main/errors"
	"main/models"
	"main/services"
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
	router.Post("/register", userController.Register)
}

func (uc *UserController) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		errors.HandleHTTPError(c, err)
	}
	user, err := uc.userService.GetUser(id)
	if err != nil {
		errors.HandleHTTPError(c, err)
	}
	return c.JSON(user)
}

func (uc *UserController) Register(c *fiber.Ctx) error {
	var user *models.User
	result, err := uc.userService.CreateUser(user)
	if err != nil {
		errors.HandleHTTPError(c, err)
	}

	return c.JSON(result)
}
