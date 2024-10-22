package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password Password `gorm:"foreignKey:UserID"`
}