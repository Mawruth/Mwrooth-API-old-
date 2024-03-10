package controllers

import (
	"github.com/gin-gonic/gin"
	"main/data/req"
	"main/models"
	"main/services"
	"main/utils"
	"net/http"
)

type StoryController struct {
	StoryService *services.StoryService
}

func NewStoryController() *StoryController {
	StoryService := services.NewStoryService()
	return &StoryController{StoryService: StoryService}
}

func SetupStoryRoutes(router *gin.RouterGroup) {
	storyController := NewStoryController()
	router.POST("/", storyController.Create)
	router.GET("/", storyController.GetAllStories)
}

func (cat *StoryController) Create(c *gin.Context) {
	var StoryReq req.StoryReq
	if err := c.Bind(&StoryReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	multipartFile, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	file, err := multipartFile.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	defer file.Close()

	filePath, err := utils.UploadImageToS3(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	Story := models.Story{
		Name:      StoryReq.Name,
		ImagePath: filePath,
	}
	result, err := cat.StoryService.Create(&Story)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, result)
}

func (cat *StoryController) GetAllStories(c *gin.Context) {
	result, err := cat.StoryService.GetAllStories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, result)
}
