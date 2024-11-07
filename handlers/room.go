package handlers

import (
	"net/http"
	"online_chat/enviroment"
	"online_chat/models"
	"online_chat/service"
	"online_chat/utils"

	"github.com/labstack/echo/v4"
)

// переделать в запись в таблице
// var (
// 	num = 0
// )

func CreateRoom(c echo.Context) error {
	token := utils.ExtractTokenFromHeaderString(c.Request().Header.Get("Authorization"))
	username := service.ExtractUsernameFromToken(token, enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET"))
	
	var (
		user models.User
		room models.Room
		room_type models.RoomType
	)

	db.Where("username = ?", username).Find(&user)

	password := c.FormValue("password")
	if password == "" {
		db.Where("type = ?", "public").Find(&room_type)
		room = models.Room{
			Name: "Рум #000",
			UserID: user.ID,
			RoomTypeID: room_type.ID,
			RoomType: room_type,
		}
	} else {
		db.Where("type = ?", "private").Find(&room_type)
		room = models.Room{
			Name: "Рум #111",
			UserID: user.ID,
			RoomTypeID: room_type.ID,
			RoomType: room_type,
			RoomPassword: models.RoomPassword{
				Password: password,
			},
		}
	}

	db.Create(&room)

	return c.JSON(http.StatusCreated, room)
}

func GetRooms(c echo.Context) error {
	var rooms []models.Room

    db.Preload("RoomPassword").Preload("RoomType").Find(&rooms)

    return c.JSON(http.StatusOK, rooms)
}

// func FindRoomByName(c echo.Context) error {
// 	room_name := c.Q
// 	var room models.Room
// 	db.Preload("RoomPassword").Preload("RoomType").Take(&room, room_name)

// 	return c.JSON(http.StatusOK, room_name)
// }

func EnterRoom(c echo.Context) error {
	var room models.Room
	db.Preload("RoomPassword").Preload("RoomType").Take(&room, c.Param("id"))
	if room.RoomType.Type == "private" {
		password := c.FormValue("password")
		if password != room.RoomPassword.Password {
			return c.String(http.StatusBadRequest, "Invalid room password")
		}
		return c.JSON(http.StatusOK, room)
	}

	return c.JSON(http.StatusOK, room)
}

func DeleteRoom(c echo.Context) error {
	var room models.Room
    db.Preload("RoomPassword").Preload("RoomType").Take(&room, c.Param("id"))
    db.Delete(&room)

    return c.String(http.StatusOK, "room deleted")
}