package controllers

import (
	"github.com/gin-gonic/gin"
	"main/data/req"
	"main/models"
	"main/services"
	"net/http"
)

type CategoryController struct {
	categoryService *services.CategoryService
}

func NewCategoryController() *CategoryController {
	categoryService := services.NewCategoryService()
	return &CategoryController{categoryService: categoryService}
}

func SetupCategoryRoutes(router *gin.RouterGroup) {
	categoryController := NewCategoryController()
	router.POST("/", categoryController.Create)
}

func (cat *CategoryController) Create(c *gin.Context) {
	var categoryReq req.CategoryReq
	if err := c.Bind(&categoryReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	// file,err := c.FormFile("image")
	// if err != nil {
	// 	return errorHandling.HandleHTTPError(c, err)
	// }
	// uniqueId := uuid.New()
	// filename := strings.Replace(uniqueId.String(), "-", "", -1)
	// fileExt := strings.Split(file.Filename, ".")[1]
	// image := fmt.Sprintf("%s.%s", filename, fileExt)
	// err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", image))
	// if err != nil {
	// 	return errorHandling.HandleHTTPError(c, err)
	// }

	category := models.Category{
		Name: categoryReq.Name,
		// ImagePath: image,
	}
	result, err := cat.categoryService.Create(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, result)
}
