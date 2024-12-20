package models

type User struct {
	ID            uint      `gorm:"primary_key"`
	Username      string
	Email         string
	Password      Password  `json:"-" gorm:"foreignKey:UserID"`
	Rooms         []Room    `gorm:"foreignKey:UserID"`
	Recovery      Recovery  `json:"-" gorm:"foreignKey:UserID"`   
}