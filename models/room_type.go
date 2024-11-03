package models

import "gorm.io/gorm"

type RoomType struct {
	gorm.Model
	Type string
}