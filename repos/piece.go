package repos

import (
	"main/config"
	"main/models"

	"gorm.io/gorm"
)

type PieceRepository struct {
	db *gorm.DB
}

func NewPieceRepository() *PieceRepository {
	db := config.GetDB()
	return &PieceRepository{db: db}
}

func (r *PieceRepository) Create(piece *models.Piece) (*models.Piece, error) {
	err := r.db.Create(piece).Error
	if err != nil {
		return nil, err
	}
	return piece, nil
}

func (r *PieceRepository) GetAllPieces() ([]models.Piece, error) {
	var pieces []models.Piece
	err := r.db.Find(&pieces).Error
	if err != nil {
		return nil, err
	}
	return pieces, nil
}

func (r *PieceRepository) GetPieceById(id int64) (*models.Piece, error) {
	var piece models.Piece
	err := r.db.Where("id = ?", id).First(&piece).Error
	if err != nil {
		return nil, err
	}
	return &piece, nil
}

func (r *PieceRepository) UpdatePiece(piece *models.Piece) (*models.Piece, error) {
	err := r.db.Save(piece).Error
	if err != nil {
		return nil, err
	}
	return piece, nil
}