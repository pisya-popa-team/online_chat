package models

type User struct {
	ID       uint     `gorm:"primary_key"`
	Username string
	Email    string
	Password Password `gorm:"foreignKey:UserID"`
	Room     []Room   `gorm:"foreignKey:UserID"`
}