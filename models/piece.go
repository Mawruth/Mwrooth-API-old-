package models

import "gorm.io/gorm"

type Piece struct {
	*gorm.Model
	Name        string       `json:"name" validate:"required" gorm:"unique"`
	Description string       `json:"description" validate:"required"`
	ARPath      string       `json:"ar_path"`
	MasterPiece bool         `json:"master_piece"`
	Images      []PieceImage `json:"images" gorm:"foreignKey:PieceID"`
	MuseumID    uint         `json:"museum_id"`
	CategoryID  uint         `json:"category_id"`
	// Category Category `json:"category" gorm:"foreignKey:CategoryID"`
}
