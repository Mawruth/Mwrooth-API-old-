package controllers

import (
	"main/errorHandling"
	"main/models"
	"main/services"

	"github.com/gofiber/fiber/v2"
)

type TypeController struct {
	typeService *services.TypeService
}

func NewTypeController() *TypeController {
	typeService := services.NewTypeService()
	return &TypeController{typeService: typeService}
}

func SetupTypeRoutes (router fiber.Router) {
	typeController := NewTypeController()
	router.Post("/", typeController.CreateType)	
	router.Get("/", typeController.GetAllTypes)
}

func (tc *TypeController) CreateType(c *fiber.Ctx) error {
	var type_ *models.Type
	if err := c.BodyParser(&type_); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	result, err := tc.typeService.Create(type_)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}

func (tc *TypeController) GetAllTypes (c *fiber.Ctx) error {
	result, err := tc.typeService.GetAllTypes()
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}