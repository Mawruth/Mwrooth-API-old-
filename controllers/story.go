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

type StoryController struct {
	StoryService *services.StoryService
}

func NewStoryController() *StoryController {
	StoryService := services.NewStoryService()
	return &StoryController{StoryService: StoryService}
}

func SetupStoryRoutes(router fiber.Router) {
	StoryController := NewStoryController()
	router.Post("/", StoryController.Create)
	router.Get("/", StoryController.GetAllStories)
}

func (cat *StoryController) Create(c *fiber.Ctx) error {
	var StoryReq req.StoryReq
	if err := c.BodyParser(&StoryReq); err != nil {
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

	Story := models.Story{
		Name: StoryReq.Name,
		ImagePath: image,
	}
	result, err := cat.StoryService.Create(&Story)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}

func (cat *StoryController) GetAllStories(c *fiber.Ctx) error {
	result, err := cat.StoryService.GetAllStories()
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}