package controllers

import (
	"github.com/gin-gonic/gin"
	"main/errorHandling"
	"main/models"
	"main/services"
	"main/utils"
	"net/http"
	"strconv"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	userService := services.NewUserService()
	return &UserController{
		userService: userService,
	}
}

func SetupUserRoutes(router *gin.RouterGroup) {
	userController := NewUserController()
	router.GET("/:id", userController.GetUser)
	router.GET("/email/:email", userController.GetUserByEmail)
	router.PATCH("/", userController.UpdateUser)
	router.POST("/register", errorHandling.ValidateRegister, userController.Register)
	router.POST("/login", userController.Login)
	router.POST("/otp/verify", userController.VerifyOTP)
	router.POST("/otp/resend", userController.ResendOTP)
}

func (uc *UserController) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	user, err := uc.userService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Register(c *gin.Context) {
	var user *models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	result, err := uc.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, result)
}

func (uc *UserController) Login(c *gin.Context) {
	var user *models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email and password are required",
		})
	}
	result, err := uc.userService.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, result)
}

func (uc *UserController) VerifyOTP(c *gin.Context) {
	var body map[string]interface{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if _, ok := body["email"].(string); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "enter a valid email",
		})
	}
	if _, ok := body["otp"].(string); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "enter a valid otp",
		})
	}
	if err := uc.userService.VerifyOTP(body["email"].(string), body["otp"].(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "otp verified",
	})
}

func (uc *UserController) ResendOTP(c *gin.Context) {
	var body map[string]interface{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := uc.userService.ResendOTP(body["email"].(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "otp resent successfully",
	})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var requestData models.UpdateUserDto

	if err := c.Bind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	user, err := uc.userService.GetUserByEmail(requestData.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	avatar, err := c.FormFile("avatar")
	if err == nil {
		avatarFile, err := avatar.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		defer avatarFile.Close()
		avatarUrl, err := utils.UploadImageToS3(avatarFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to upload image",
			})
		}
		user.Avatar = avatarUrl
	}

	if requestData.FullName != "" {
		user.FullName = requestData.FullName
	}

	if requestData.Password != "" {
		user.Password = requestData.Password
	}

	result, err := uc.userService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, result)
}

func (uc *UserController) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := uc.userService.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, user)
}
