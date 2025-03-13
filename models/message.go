package models

import "time"

type Message struct {
	ID           uint        `json:"id" gorm:"primary_key"`
	MessageType  MessageType `json:"message_type"`
	Content      string      `json:"content"`
	SentAt       time.Time	 `json:"sent_at"`
	RoomID       uint        `json:"room_id"`
	UserID       uint        `json:"-"`
}