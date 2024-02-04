package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the configuration parameters
type Config struct {
	PORT	string
}

// LoadConfig loads the configuration from the .env file
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Error loading .env file: %w", err)
	}

	fmt.Println("Port config")
	fmt.Println(os.Getenv("PORT"))
	config := &Config{
		PORT: os.Getenv("PORT"),
	}

	return config, nil
}