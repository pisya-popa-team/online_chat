package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name         string
	UserID       uint
	RoomTypeID   uint
	RoomType     RoomType
	RoomPassword RoomPassword `gorm:"foreignKey:RoomID"`
}