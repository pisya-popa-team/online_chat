package models

type Message struct {
	ID           uint        `json:"-" gorm:"primary_key"`
	MessageType  MessageType 
	Content      string      
	SentAt       string	 
	RoomID       uint        
	UserID       uint
	User         User       `gorm:"foreignKey:UserID"`
}

type UserMessage struct {
	ID	     uint	`json:"-"`
	Message  Message
	Username string
}