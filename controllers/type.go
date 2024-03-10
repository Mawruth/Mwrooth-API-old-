package controllers

import (
	"github.com/gin-gonic/gin"
	"main/data/req"
	"main/models"
	"main/services"
	"main/utils"
	"net/http"
)

type TypeController struct {
	typeService *services.TypeService
}

func NewTypeController() *TypeController {
	typeService := services.NewTypeService()
	return &TypeController{typeService: typeService}
}

func SetupTypeRoutes(router *gin.RouterGroup) {
	typeController := NewTypeController()
	router.POST("/", typeController.CreateType)
	router.GET("/", typeController.GetAllTypes)
}

func (tc *TypeController) CreateType(c *gin.Context) {
	var (
		typeReq req.Type
		type_   *models.Type = &models.Type{}
	)
	if err := c.Bind(&typeReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	image, err := c.FormFile("image")
	if err == nil {
		imageFile, err := image.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		defer imageFile.Close()
		imageUrl, err := utils.UploadImageToS3(imageFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to upload image",
			})
		}
		type_.Image = imageUrl
	}
	type_.Name = typeReq.Name
	result, err := tc.typeService.Create(type_)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, result)
}

func (tc *TypeController) GetAllTypes(c *gin.Context) {
	result, err := tc.typeService.GetAllTypes()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, result)
}
