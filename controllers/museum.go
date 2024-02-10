package controllers

import (
	"fmt"
	"main/data/req"
	"main/errorHandling"
	"main/services"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/google/uuid"
)

type MuseumController struct {
	museumService *services.MuseumService
}

func NewMuseumController() *MuseumController {
	museumService := services.NewMuseumService()
	return &MuseumController{museumService: museumService}
}

func SetupMuseumRoutes(router fiber.Router) {
	museumController := NewMuseumController()
	router.Post("/", museumController.Create)
	router.Get("/", museumController.GetAll)
	router.Get("/:id", museumController.GetByID)
}

func (m *MuseumController) Create(c *fiber.Ctx) error {

	var museumReq req.MuseumReq
	if err := c.BodyParser(&museumReq); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	files := form.File["images"]
	var images []string

	for _, file := range files {
		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := strings.Split(file.Filename, ".")[1]
		image := fmt.Sprintf("%s.%s", filename, fileExt)
		err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", image))
		if err != nil {
			return errorHandling.HandleHTTPError(c, err)
		}
		images = append(images, image)
		museumReq.Images = images
	}
	result, _ := m.museumService.CreateMuseum(museumReq)

	return c.JSON(result)
}

func (m *MuseumController) GetAll(c *fiber.Ctx) error {
	ratingParam := utils.CopyString(c.Query("rating"))
	typesParam := utils.CopyString(c.Query("types"))
	cityParam := utils.CopyString(c.Query("city"))

	if ratingParam != "" {
		museums, err := m.museumService.GetByRating(ratingParam)
		if err != nil {
			return errorHandling.HandleHTTPError(c, err)
		}
		return c.Status(fiber.StatusOK).JSON(museums)
	}

	if typesParam != "" {
		museums, err := m.museumService.GetByTypes(typesParam)
		if err != nil {
			return errorHandling.HandleHTTPError(c, err)
		}
		return c.Status(fiber.StatusOK).JSON(museums)
	}

	if cityParam != "" {
		museums, err := m.museumService.GetByCity(cityParam)
		if err != nil {
			return errorHandling.HandleHTTPError(c, err)
		}
		return c.Status(fiber.StatusOK).JSON(museums)
	}
	museums, err := m.museumService.GetAll()
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(museums)
}

func (m *MuseumController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	museum, err := m.museumService.GetByID(id)
	return c.Status(fiber.StatusOK).JSON(museum)
}
