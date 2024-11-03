package database

import (
	"online_chat/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func GetDBConnection() *gorm.DB {
	if dbConnection == nil {
		connectDB()
	}

	return dbConnection
}

func initRecords(db *gorm.DB) {
	room_type := []*models.RoomType{
		{Type: "public"},
        {Type: "private"},
	}
	_ = db.Create(room_type)
}

func connectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	_ = db.AutoMigrate(
		&models.User{},
		&models.Password{},
		&models.Room{},
		&models.RoomType{},
		&models.RoomPassword{},
	)

	initRecords(db)
	
	dbConnection = db

	return dbConnection
}