package main

import (
	"fmt"
	"log"
	"main/config"
	"main/controllers"
	"main/middlewares"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./uploads")
	apiGroup := app.Group("/api/v1")
	apiGroup.Use(middlewares.CheckAccessToken)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	loadConfig, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading loadConfig:", err)
	}
	PORT := loadConfig.API_PORT

	if err := loadConfig.DB.AutoMigrate(
		&models.User{}, &models.Category{}, &models.Type{}, &models.Museum{}, &models.Piece{}, &models.PieceImages{}, &models.MuseumImages{}, &models.Story{}, &models.Review{},
	); err != nil {
		log.Fatalf("Error running migrations: %s", err.Error())
	}
	controllers.SetupUserRoutes(apiGroup.Group("users"))
	controllers.SetupTypeRoutes(apiGroup.Group("types"))
	controllers.SetupMuseumRoutes(apiGroup.Group("museums"))
	controllers.SetupCategoryRoutes(apiGroup.Group("categories"))
	controllers.SetupStoryRoutes(apiGroup.Group("stories"))
	controllers.SetupPieceRoute(apiGroup.Group("pieces"))
	controllers.SetupReviewRoutes(apiGroup.Group("reviews"))
	log.Fatalf("Error running server: %s\n", app.Listen(PORT))
}
