package models

type Message struct {
	ID           uint        `json:"-" gorm:"primary_key"`
	MessageType  MessageType 
	Content      string      
	SentAt       string	 
	RoomID       uint        
	UserID       uint
}