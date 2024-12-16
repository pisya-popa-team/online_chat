package models

import "time"

type Message struct {
	ID           uint        `gorm:"primary_key"`
	MessageType  MessageType
	Content      string
	SentAt       time.Time
	RoomID       uint
	UserID       uint
}