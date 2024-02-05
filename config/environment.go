package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

// Config holds the configuration parameters
type Config struct {
	API_PORT string
	DB       *gorm.DB
}

// LoadConfig loads the configuration from the .env file
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Error loading .env file: %w", err)
	}

	port := os.Getenv("API_PORT")
	if port == "" {
		port = ":3000"
	}

	config := &Config{
		API_PORT: port,
	}

	var err error
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		db, err = gorm.Open(postgres.Open(dsn))
		if err != nil {
			err = fmt.Errorf("Error connecting to database: %w", err)
		}
	})
	if err != nil {
		return nil, err
	}

	config.DB = db

	return config, nil
}

func GetDB() *gorm.DB {
	return db
}
