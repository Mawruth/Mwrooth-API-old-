package controllers

import (
	"fmt"
	"main/data/req"
	"main/errorHandling"
	"main/models"
	"main/services"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CategoryController struct {
	categoryService *services.CategoryService
}

func NewCategoryController() *CategoryController {
	categoryService := services.NewCategoryService()
	return &CategoryController{categoryService: categoryService}
}

func SetupCategoryRoutes(router fiber.Router) {
	categoryController := NewCategoryController()
	router.Post("/", categoryController.Create)
}

func (cat *CategoryController) Create(c *fiber.Ctx) error {
	var categoryReq req.CategoryReq
	if err := c.BodyParser(&categoryReq); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	file,err := c.FormFile("image")
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", image))
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	category := models.Category{
		Name: categoryReq.Name,
		ImagePath: image,
	}
	result, err := cat.categoryService.Create(&category)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}