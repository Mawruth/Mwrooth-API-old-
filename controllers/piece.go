package controllers

import (
	"main/data/req"
	"main/errorHandling"
	"main/models"
	"main/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
	route.Get("/", pieceController.GetAll)
	route.Get("/master-piece/:id", pieceController.MakeMasterPiece)
	route.Get("/:id", pieceController.GetById)
}

func (pieceController *PieceController) Create(c *fiber.Ctx) error {
	var piece req.PieceReq
	if err := c.BodyParser(&piece); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}

	// fmt.Println(piece)

	files := form.File["images"]
	var images []string

	for _, file := range files {
		// uniqueId := uuid.New()
		// filename := strings.Replace(uniqueId.String(), "-", "", -1)
		// fileExt := strings.Split(file.Filename, ".")[1]
		// image := fmt.Sprintf("%s.%s", filename, fileExt)
		// err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", image))
		// images = append(images, image)
		fileData, err := file.Open()
		if err != nil {
			return errorHandling.HandleHTTPError(c, err)
		}
		defer fileData.Close()
		fileUrl, err := uploadImageToS3(fileData)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "Failed to upload image",
			})
		}
		images = append(images, fileUrl)
	}

	// ar_file := form.File["ar_obj"]
	// var ar_path string
	// for _,file := range ar_file {
	// 	uniqueId := uuid.New()
	// 	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	// 	fileExt := strings.Split(file.Filename, ".")[1]
	// 	path := fmt.Sprintf("%s.%s", filename, fileExt)
	// 	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", path))
	// 	ar_path = path
	// 	if err != nil {
	// 		return errorHandling.HandleHTTPError(c, err)
	// 	}
	// }

	var pieceImages []models.PieceImage

	for _, img := range images {
		piceImage := models.PieceImage{
			ImagePath: img,
		}
		pieceImages = append(pieceImages, piceImage)
	}

	newPice := models.Piece{
		Name:        piece.Name,
		Description: piece.Description,
		MasterPiece: piece.MasterPiece,
		CategoryID:  piece.CategoryID,
		MuseumID:    piece.MuseumID,
		Images:      pieceImages,
		ARPath:      piece.ARPath,
	}
	result, err := pieceController.pieceService.CreatePiece(newPice)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(result)
}

func (pieceController *PieceController) GetAll(c *fiber.Ctx) error {
	pieces, err := pieceController.pieceService.GetAllPieces()
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(pieces)
}

func (pieceController *PieceController) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	piece, err := pieceController.pieceService.GetPieceById(idInt)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(piece)
}

func (pieceController *PieceController) MakeMasterPiece(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	piece, err := pieceController.pieceService.GetPieceById(idInt)
	if err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	piece.MasterPiece = true
	if _, err := pieceController.pieceService.UpdatePiece(&piece); err != nil {
		return errorHandling.HandleHTTPError(c, err)
	}
	return c.JSON(piece)
}
