package models

import "time"

type Message struct {
	ID           uint        `json:"-" gorm:"primary_key"`
	MessageType  MessageType 
	Content      string      
	SentAt       time.Time	 
	RoomID       uint        
	UserID       uint        `json:"-"`
	User         User        `gorm:"foreignKey:UserID"`
}