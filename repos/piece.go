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