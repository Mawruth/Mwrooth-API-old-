package controllers

import (
	"main/errorHandling"
	"main/models"
	"main/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ReviewController struct {
	ReviewService *services.ReviewService
}

func NewReviewController() *ReviewController {
	ReviewService := services.NewReviewService()
	return &ReviewController{ReviewService: ReviewService}
}

func SetupReviewRoutes(router fiber.Router) {
	ReviewController := NewReviewController()
	router.Post("/", ReviewController.Create)
	router.Get("/", ReviewController.GetAllReviews)
	router.Get("/museum/:id", ReviewController.GetReviewByMuseum)
	router.Put("/:id", ReviewController.Update)
}

func (cat *ReviewController) Create(c *fiber.Ctx) error {
	var ReviewReq models.Review
	if err := c.BodyParser(&ReviewReq); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	result, err := cat.ReviewService.Create(&ReviewReq)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}

func (cat *ReviewController) GetAllReviews(c *fiber.Ctx) error {
	result, err := cat.ReviewService.GetAllReviews()
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}

func (cat *ReviewController) GetReviewByMuseum(c *fiber.Ctx) error {
	id := c.Params("id")
	museumId, err := strconv.Atoi(id)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	result, err := cat.ReviewService.GetReviewByMuseum(museumId)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}

func (cat *ReviewController) Update(c *fiber.Ctx) error {
	var Review *models.Review
	if err := c.BodyParser(&Review); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	result, err := cat.ReviewService.Update(Review)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}
