package models

import "gorm.io/gorm"

type Story struct {
	*gorm.Model
	Name string `json:"name" validate:"required" gorm:"unique"`
	Description string `json:"description" validate:"required"`
	ImagePath string `json:"image_path" validate:"required"`
	MuseumID uint `json:"museum_id" validate:"required"`
}

