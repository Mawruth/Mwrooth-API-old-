package services

import (
	"main/models"
	"main/repos"
)

type PieceService struct {
	pieceRepository *repos.PieceRepository
}

func NewPieceService() *PieceService {
	return &PieceService{
		pieceRepository: repos.NewPieceRepository(),
	}
}	

func (s *PieceService) CreatePiece(piece models.Piece) (*models.Piece, error) {
	return s.pieceRepository.Create(&piece)
}

func (s *PieceService) GetAllPieces() ([]models.Piece, error) {
	return s.pieceRepository.GetAll()
}

func (s *PieceService) GetPieceById(id int) (models.Piece, error) {
	return s.pieceRepository.GetById(id)
}


func (s *PieceService) UpdatePiece(piece *models.Piece) (*models.Piece, error) {
	return s.pieceRepository.Update(piece)
}