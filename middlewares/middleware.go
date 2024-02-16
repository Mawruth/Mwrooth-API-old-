package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

var ExceptionEndpoints = []string{
	"/api/v1/users",
	"/api/v1/museums",
	"/api/v1/pieces",
	"/api/v1/reviews",
}

func CheckAccessToken(c *fiber.Ctx) error {
	for _, exc := range ExceptionEndpoints {
		if strings.Contains(c.Path(), exc) {
			return c.Next()
		}
	}
	token := c.Get("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	if err := verifyToken(token); err != nil {
		fmt.Println(token)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Next()
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
