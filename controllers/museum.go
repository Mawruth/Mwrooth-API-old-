package controllers

import (
	"github.com/gin-gonic/gin"
	"main/data/req"
	"main/services"
	"main/utils"
	"net/http"
	"strconv"
)

type MuseumController struct {
	museumService *services.MuseumService
}

func NewMuseumController() *MuseumController {
	museumService := services.NewMuseumService()
	return &MuseumController{museumService: museumService}
}

func SetupMuseumRoutes(router *gin.RouterGroup) {
	museumController := NewMuseumController()
	router.POST("/", museumController.Create)
	router.GET("/", museumController.GetAll)
	router.GET("/:id", museumController.GetByID)
}

func (m *MuseumController) Create(c *gin.Context) {
	var museumReq req.MuseumReq
	if err := c.Bind(&museumReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	files := form.File["images"]
	var images []string

	for _, file := range files {
		fileData, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		defer fileData.Close()
		fileUrl, err := utils.UploadImageToS3(fileData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to upload image",
			})
		}
		images = append(images, fileUrl)
	}
	museumReq.Images = images
	result, _ := m.museumService.CreateMuseum(museumReq)

	c.JSON(http.StatusOK, result)
}

func (m *MuseumController) GetAll(c *gin.Context) {
	ratingParam := c.Query("rating")
	typesParam := c.Query("types")
	cityParam := c.Query("city")

	if ratingParam != "" {
		museums, err := m.museumService.GetByRating(ratingParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, museums)
	}

	if typesParam != "" {
		museums, err := m.museumService.GetByTypes(typesParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, museums)
	}

	if cityParam != "" {
		museums, err := m.museumService.GetByCity(cityParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, museums)
	}
	museums, err := m.museumService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, museums)
}

func (m *MuseumController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	museum, err := m.museumService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, museum)
}
