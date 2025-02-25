package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	FullName string  `json:"full_name" `
	UserName string  `json:"user_name" validate:"required" gorm:"unique"`
	Email    string  `json:"email" validate:"required" gorm:"unique"`
	Password string  `json:"password" validate:"required"`
	OTP      *string `json:"-" gorm:"otp;null;default:null"`
	Avatar   string  `json:"avatar"`
}

type UpdateUserDto struct {
	FullName string `form:"full_name"`
	UserName string `form:"user_name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Avatar   string `form:"avatar"`
}
