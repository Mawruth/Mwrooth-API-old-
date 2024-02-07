package main

import (
	"fmt"
	"log"
	"main/config"
	"main/controllers"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./uploads")
	apiGroup := app.Group("/api/v1")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
	}
	PORT := config.API_PORT

	if err := config.DB.AutoMigrate(
		&models.User{},&models.Category{},&models.Type{},&models.Museum{},&models.Piece{},&models.PieceImages{},
	); err != nil {
		log.Fatalf("Error running migrations: %s", err.Error())
	}
	controllers.SetupUserRoutes(apiGroup.Group("users"))
	controllers.SetupTypeRoutes(apiGroup.Group("types"))
	controllers.SetupMuseumRoutes(apiGroup.Group("museums"))
	controllers.SetupCategoryRoutes(apiGroup.Group("categories"))
	controllers.SetupPieceRoute(apiGroup.Group("pieces"))
	app.Listen(PORT)
}
