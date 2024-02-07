package models

import "gorm.io/gorm"

type Category struct {
	*gorm.Model
	Name 	 	string  `json:"name" validate:"required" gorm:"unique"`
	ImagePath 	string  `json:"image_path" validate:"required"`
	Pieces 		[]Piece `json:"pieces" gorm:"foreignKey:CategoryID"`
}