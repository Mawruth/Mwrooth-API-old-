package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/config"
	"main/controllers"
	"main/models"
)

func main() {
	loadConfig, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading loadConfig:", err)
	}
	PORT := loadConfig.API_PORT

	fmt.Println("Migrating database...")
	if err := loadConfig.DB.AutoMigrate(
		&models.User{}, &models.Category{}, &models.Type{}, &models.Museum{}, &models.Piece{}, &models.PieceImage{}, &models.MuseumImage{}, &models.Story{}, &models.Review{},
	); err != nil {
		log.Fatalf("Error running migrations: %s", err.Error())
	} else {
		fmt.Println("Migrations successful")
	}

	app := gin.Default()
	apiGroup := app.Group("/api/v1")
	controllers.SetupUserRoutes(apiGroup.Group("users"))
	controllers.SetupTypeRoutes(apiGroup.Group("types"))
	controllers.SetupMuseumRoutes(apiGroup.Group("museums"))
	controllers.SetupCategoryRoutes(apiGroup.Group("categories"))
	controllers.SetupStoryRoutes(apiGroup.Group("stories"))
	controllers.SetupPieceRoute(apiGroup.Group("pieces"))
	controllers.SetupReviewRoutes(apiGroup.Group("reviews"))
	// apiGroup.Use(middlewares.CheckAccessToken)

	err = app.Run(PORT)
	if err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}
}
