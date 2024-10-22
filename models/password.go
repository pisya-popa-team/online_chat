package models

import "gorm.io/gorm"

type Password struct {
	gorm.Model
	Hash   string
	UserID uint
}