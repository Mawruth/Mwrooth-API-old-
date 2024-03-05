package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserId   uint    `json:"user_id" validate:"required"`
	MuseumID uint    `json:"museum_id" validate:"required"`
	Content  string  `json:"content" validate:"required"`
	Creator  *User   `json:"creator" gorm:"-"`
	Rating   float32 `json:"rating" validate:"required"`
}
