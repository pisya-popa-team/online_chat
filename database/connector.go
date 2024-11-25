package database

import (
	"online_chat/enviroment"
	"online_chat/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func GetDBConnection() *gorm.DB {
	if dbConnection == nil {
		node_env := enviroment.GoDotEnvVariable("NODE_ENV")
		if node_env == "development"{
			connectDBSqlite()
		} else if node_env == "production"{
			connectDBPostgres()
		} else {
			panic("undefined db connection")
		}
	}

	return dbConnection
}

func connectDB(db *gorm.DB, err error) *gorm.DB {
	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(
		&models.User{},
		&models.Password{},
		&models.Room{},
		&models.RoomPassword{},
		&models.Recovery{},
	)
	
	dbConnection = db

	return dbConnection
}

func connectDBSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	return connectDB(db, err)
}

func connectDBPostgres() *gorm.DB {
	db_url := enviroment.GoDotEnvVariable("DB_CONNECTION")
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})

	return connectDB(db, err)
}