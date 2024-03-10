package models

import "gorm.io/gorm"

type MuseumImage struct {
	*gorm.Model
	MuseumID  uint
	ImagePath string `json:"image_path" validate:"required"`
}
