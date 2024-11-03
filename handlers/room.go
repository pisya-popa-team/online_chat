package handlers

import (
	"net/http"
	"online_chat/enviroment"
	"online_chat/models"
	"online_chat/service"
	"online_chat/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

//переделать в запись в таблице
var (
	num = 0
)

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
			Name: "Рум #" + strconv.Itoa(num),
			UserID: user.ID,
			RoomTypeID: room_type.ID,
			RoomType: room_type,
		}
	} else {
		db.Where("type = ?", "private").Find(&room_type)
		room = models.Room{
			Name: "Рум #" + strconv.Itoa(num),
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
// 	var room models.Room
// 	db.Preload("RoomPassword").Preload("RoomType").Take(&room, c.Param("name"))

// 	if room.Name == "" {
// 		return c.String(http.StatusNotFound, "there's no such room under that name")
// 	}

// 	return c.JSON(http.StatusOK, room)
// }

func DeleteRoom(c echo.Context) error {
	var room models.Room
    db.Preload("RoomPassword").Preload("RoomType").Take(&room, c.Param("id"))
    db.Delete(&room)

    return c.String(http.StatusOK, "room deleted")
}