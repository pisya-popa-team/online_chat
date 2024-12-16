package models

type Room struct {
	ID           uint         `gorm:"primary_key"`
	Name         string       `default:""`
	UserID       uint
	RoomType     RoomType     
	RoomPassword RoomPassword `json:"-" gorm:"foreignKey:RoomID"`
	Users        []User       `json:"-" gorm:"many2many:room_users"`
	Messages     []Message    `json:"-" gorm:"foreignKey:RoomID"`
}