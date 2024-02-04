package main

import (
	"fmt"
	"main/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	config,err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
	}

	PORT := config.PORT
	if (PORT == "") {
		PORT = ":3000"
	}
	fmt.Println("Server is running on port " + PORT)
	app.Listen(PORT)
}