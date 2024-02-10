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

func (r *PieceRepository) GetAll() ([]models.Piece, error) {
	var pieces []models.Piece
	err := r.db.Find(&pieces).Error
	if err != nil {
		return nil, err
	}
	return pieces, nil
}

func (r *PieceRepository) GetById(id int) (models.Piece, error) {
	var piece models.Piece
	err := r.db.First(&piece, id).Error
	if err != nil {
		return models.Piece{}, err
	}
	return piece, nil
}

func (r *PieceRepository) Update(piece *models.Piece) (*models.Piece, error) {
	err := r.db.Save(piece).Error
	if err != nil {
		return nil, err
	}
	return piece, nil
}
