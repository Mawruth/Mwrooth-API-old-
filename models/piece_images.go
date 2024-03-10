package models

import "gorm.io/gorm"

type PieceImage struct {
	*gorm.Model
	PieceID   uint
	ImagePath string `json:"image_path" validate:"required"`
}
