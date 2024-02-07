package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	FullName    string `json:"full_name" `
	UserName    string `json:"user_name" validate:"required" gorm:"unique"`
	Email       string `json:"email" validate:"required" gorm:"unique"`
	PhoneNumber string `json:"phone_number" validate:"required" gorm:"unique"`
	Password    string `json:"password" validate:"required"`
}
