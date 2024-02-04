package errorHandling

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"main/models"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

var Validator = validator.New()

func HandleHTTPError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func ValidateRegister(c *fiber.Ctx) error {
	var errors []*IError
	body := new(models.User)
	c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}
