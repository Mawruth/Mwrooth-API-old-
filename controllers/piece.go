package controllers

import (
	"github.com/gin-gonic/gin"
	"main/data/req"
	"main/models"
	"main/services"
	"main/utils"
	"net/http"
	"strconv"
)

type PieceController struct {
	pieceService *services.PieceService
}

func NewPieceController() *PieceController {
	pieceService := services.NewPieceService()
	return &PieceController{pieceService: pieceService}
}

func SetupPieceRoute(route *gin.RouterGroup) {
	pieceController := NewPieceController()
	route.POST("/", pieceController.Create)
	route.GET("/", pieceController.GetAll)
	route.GET("/master-piece/:id", pieceController.MakeMasterPiece)
	route.GET("/:id", pieceController.GetById)
}

func (pieceController *PieceController) Create(c *gin.Context) {
	var piece req.PieceReq
	if err := c.Bind(&piece); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	files := form.File["images"]
	var images []string

	for _, file := range files {
		fileData, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		defer fileData.Close()
		fileUrl, err := utils.UploadImageToS3(fileData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to upload image",
			})
		}
		images = append(images, fileUrl)
	}

	var pieceImages []models.PieceImage

	for _, img := range images {
		pieceImage := models.PieceImage{
			ImagePath: img,
		}
		pieceImages = append(pieceImages, pieceImage)
	}

	newPiece := models.Piece{
		Name:        piece.Name,
		Description: piece.Description,
		MasterPiece: piece.MasterPiece,
		CategoryID:  piece.CategoryID,
		MuseumID:    piece.MuseumID,
		Images:      pieceImages,
		ARPath:      piece.ARPath,
	}
	result, err := pieceController.pieceService.CreatePiece(newPiece)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, result)
}

func (pieceController *PieceController) GetAll(c *gin.Context) {
	pieces, err := pieceController.pieceService.GetAllPieces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, pieces)
}

func (pieceController *PieceController) GetById(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	piece, err := pieceController.pieceService.GetPieceById(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, piece)
}

func (pieceController *PieceController) MakeMasterPiece(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	piece, err := pieceController.pieceService.GetPieceById(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	piece.MasterPiece = true
	if _, err := pieceController.pieceService.UpdatePiece(&piece); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, piece)
}
