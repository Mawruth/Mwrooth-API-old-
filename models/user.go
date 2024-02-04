package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	FullName string `json:"full_name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
