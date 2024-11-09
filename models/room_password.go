package models

type RoomPassword struct {
	ID       uint   `gorm:"primary_key"`
	Password string `default:""`
	RoomID   uint   
}