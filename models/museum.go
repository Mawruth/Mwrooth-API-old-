package models

import "gorm.io/gorm"

type Museum struct {
	*gorm.Model
	Name        string        `json:"name" validate:"required" gorm:"unique"`
	Description string        `json:"description" validate:"required"`
	WorkTime    string        `json:"work_time" validate:"required"`
	Country     string        `json:"country" validate:"required"`
	City        string        `json:"city" validate:"required"`
	Street      string        `json:"street" validate:"required"`
	Rating      float32       `json:"rating"`
	Types       []Type        `json:"types" validate:"required" gorm:"many2many:museum_types;"`
	Pieces      []Piece       `json:"pieces" gorm:"foreignKey:MuseumID"`
	Images      []MuseumImage `json:"images" gorm:"foreignKey:MuseumID"`
	Reviews     []Review      `json:"reviews" gorm:"foreignKey:MuseumID"`
}
