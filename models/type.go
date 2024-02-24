package models

import "gorm.io/gorm"

type Type struct {
	*gorm.Model
	Name    string    `json:"name" validate:"required" gorm:"unique"`
	Museums []*Museum `json:"museums" gorm:"many2many:museum_types;"`
	Image   string    `json:"image"`
}
