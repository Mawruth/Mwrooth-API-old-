package errorHandling

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"main/models"
	"net/http"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

var Validator = validator.New()

func ValidateRegister(c *gin.Context) {
	var errors []*IError
	body := new(models.User)
	c.Bind(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		c.JSON(http.StatusBadRequest, errors)
	}
	c.Next()
}
