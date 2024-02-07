package models

import "gorm.io/gorm"

type MuseumImages struct {
	*gorm.Model
	MuseumID   uint
	Image_path string `json:"image_path" validate:"required"`
}
