package models

import "gorm.io/gorm"

type PieceImages struct {
	*gorm.Model
	PieceID uint
	Image_path string `json:"image_path" validate:"required"`
}