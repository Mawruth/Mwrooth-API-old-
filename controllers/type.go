package controllers

import (
	"main/data/req"
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

func SetupTypeRoutes(router fiber.Router) {
	typeController := NewTypeController()
	router.Post("/", typeController.CreateType)
	router.Get("/", typeController.GetAllTypes)
}

func (tc *TypeController) CreateType(c *fiber.Ctx) error {
	var (
		typeReq req.Type
		type_   *models.Type = &models.Type{}
	)
	if err := c.BodyParser(&typeReq); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	image, err := c.FormFile("image")
	if err == nil {
		imageFile, err := image.Open()
		if err != nil {
			return errorHandling.HandleHTTPError(c, err)
		}
		defer imageFile.Close()
		imageUrl, err := uploadImageToS3(imageFile)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "Failed to upload image",
			})
		}
		type_.Image = imageUrl
	}
	type_.Name = typeReq.Name
	result, err := tc.typeService.Create(type_)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}

func (tc *TypeController) GetAllTypes(c *fiber.Ctx) error {
	result, err := tc.typeService.GetAllTypes()
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}
