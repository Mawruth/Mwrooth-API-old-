package controllers

import (
	"github.com/gin-gonic/gin"
	"main/models"
	"main/services"
	"net/http"
	"strconv"
)

type ReviewController struct {
	ReviewService *services.ReviewService
}

func NewReviewController() *ReviewController {
	ReviewService := services.NewReviewService()
	return &ReviewController{ReviewService: ReviewService}
}

func SetupReviewRoutes(router *gin.RouterGroup) {
	reviewController := NewReviewController()
	router.POST("/", reviewController.Create)
	router.GET("/", reviewController.GetAllReviews)
	router.GET("/museum/:id", reviewController.GetReviewByMuseum)
	router.PUT("/:id", reviewController.Update)
}

func (cat *ReviewController) Create(c *gin.Context) {
	var ReviewReq models.Review
	if err := c.Bind(&ReviewReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	result, err := cat.ReviewService.Create(&ReviewReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, result)
}

func (cat *ReviewController) GetAllReviews(c *gin.Context) {
	result, err := cat.ReviewService.GetAllReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, result)
}

func (cat *ReviewController) GetReviewByMuseum(c *gin.Context) {
	museumId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	result, err := cat.ReviewService.GetReviewByMuseum(museumId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, result)
}

func (cat *ReviewController) Update(c *gin.Context) {
	var Review *models.Review
	if err := c.Bind(&Review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	result, err := cat.ReviewService.Update(Review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, result)
}
