package models

import "time"

type Recovery struct {
	ID        uint      `gorm:"primary_key"`
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
	IsUsed    bool      `gorm:"default:false"`
	UserID    uint
}