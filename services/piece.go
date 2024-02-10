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
	return s.pieceRepository.GetAllPieces()
}

func (s *PieceService) GetPieceById(id int64) (*models.Piece, error) {
	return s.pieceRepository.GetPieceById(id)
}

func (s *PieceService) UpdatePiece(piece *models.Piece) (*models.Piece, error) {
	return s.pieceRepository.UpdatePiece(piece)
}