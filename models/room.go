package models

type Room struct {
	ID           uint         `gorm:"primary_key"`
	Name         string       `default:""`
	UserID       uint
	RoomTypeID   uint
	RoomType     RoomType
	RoomPassword RoomPassword `gorm:"foreignKey:RoomID"`
}