package controllers

import (
	"fmt"
	"main/data/req"
	"main/errorHandling"
	"main/services"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MuseumController struct {
	museumService *services.MuseumService
}

func NewMuseumController() *MuseumController {
	museumService := services.NewMuseumService()
	return &MuseumController{museumService: museumService}
}

func SetupMuseumRoutes (router fiber.Router) {
	museumController := NewMuseumController()
	router.Post("/", museumController.Create)
}

func (m *MuseumController) Create(c *fiber.Ctx) error {
	
	var museumReq req.MuseumReq
	if err := c.BodyParser(&museumReq); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	} 

	form, err := c.MultipartForm();
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	files := form.File["images"]
	var images []string

	for _,file := range files {
		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := strings.Split(file.Filename, ".")[1]
		image := fmt.Sprintf("%s.%s", filename, fileExt)
		err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", image))
		images = append(images, image)
		if err != nil {
			return errorHandling.HandleHTTPError(c, err)
		}

	}
	result,_ := m.museumService.CreateMuseum(museumReq)

	return c.JSON(result)
}