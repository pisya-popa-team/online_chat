package models

type Type string

const (
	Private Type = "private"
	Public  Type = "public"
)
type RoomType struct {
	ID   uint     `gorm:"primary_key"`
	Type Type
}