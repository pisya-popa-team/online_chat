package models

import "gorm.io/gorm"

type RoomPassword struct {
	gorm.Model
	Password string `default:""`
	RoomID   uint
}