package models

type Room struct {
	ID           uint         `gorm:"primary_key"`
	Name         string       `default:""`
	UserID       uint
	RoomType     RoomType
	RoomPassword RoomPassword `gorm:"foreignKey:RoomID"`
}