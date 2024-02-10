package controllers

import (
	"fmt"
	"main/data/req"
	"main/errorHandling"
	"main/models"
	"main/services"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PieceController struct {
	pieceService *services.PieceService
}

func NewPieceController() *PieceController {
	pieceService := services.NewPieceService()
	return &PieceController{pieceService: pieceService}
}

func SetupPieceRoute(route fiber.Router) {
	pieceController := NewPieceController()
	route.Post("/", pieceController.Create)
	route.Get("/", pieceController.GetAllPieces)
	route.Get("/make-master/:id", pieceController.MakeMasterPiece)
	route.Get("/:id", pieceController.GetPieceById)
}

func (pieceController *PieceController) Create(c *fiber.Ctx) error {
	var piece req.PieceReq
	if err := c.BodyParser(&piece); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	form, err := c.MultipartForm();
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	fmt.Println(piece)

	files := form.File["images"]
	var images []string

	for _,file := range files {
		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := strings.Split(file.Filename, ".")[1]
		image := fmt.Sprintf("%s.%s", filename, fileExt)
		err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", image))
		images = append(images, image)
		if err != nil {
			return errorHandling.HandleHTTPError(c, err)
		}
	}

	ar_file := form.File["ar_obj"]
	var ar_path string
	for _,file := range ar_file {
		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := strings.Split(file.Filename, ".")[1]
		path := fmt.Sprintf("%s.%s", filename, fileExt)
		err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", path))
		ar_path = path
		if err != nil {
			return errorHandling.HandleHTTPError(c, err)
		}
	}
	
	var pieceImages []models.PieceImages

	for _,img := range images {
		piceImage := models.PieceImages{
			Image_path: img,
		}
		pieceImages = append(pieceImages, piceImage)
	}


	newPice := models.Piece{
		Name: piece.Name,
		Description: piece.Description,
		Master_piece: piece.Master_piece,
		CategoryID: piece.CategoryID,
		MuseumID: piece.MuseumID,
		Images: pieceImages,
		AR_Path: ar_path,
	}
	result, err := pieceController.pieceService.CreatePiece(newPice)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}

func (pieceController *PieceController) GetAllPieces(c *fiber.Ctx) error {
	result, err := pieceController.pieceService.GetAllPieces()
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}

 func (pieceController *PieceController) MakeMasterPiece(c *fiber.Ctx) error {
 	id := c.Params("id")
 	p_id, err := strconv.ParseInt(id, 10, 64)
 	if err != nil {
 		return errorHandling.HandleHTTPError(c, err)
 	}
 	p, err := pieceController.pieceService.GetPieceById(p_id)
 	if err != nil {
 		return errorHandling.HandleHTTPError(c, err)
 	}
 	p.Master_piece = true

 	p,err = pieceController.pieceService.UpdatePiece(p)
 	if err != nil {
 		return errorHandling.HandleHTTPError(c, err)
 	}
 	return c.JSON(p)
 }

func (pieceController *PieceController) GetPieceById(c *fiber.Ctx) error {
	id := c.Params("id")
	p_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	p, err := pieceController.pieceService.GetPieceById(p_id)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(p)
}
