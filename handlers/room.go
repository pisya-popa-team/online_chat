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



func CreateRoom(c echo.Context) error {
	token := utils.ExtractTokenFromHeaderString(c.Request().Header.Get("Authorization"))
	username := service.ExtractUsernameFromToken(token, enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET"))
	
	var (
		user models.User
		room models.Room
	)

	db.Where("username = ?", username).Find(&user)

	password := c.FormValue("password")
	if password == "" {
		room = models.Room{
			UserID: user.ID,
			RoomType: models.RoomType{
				Type: models.Public,
			},
		}

	} else {
		room = models.Room{
			UserID: user.ID,
			RoomType: models.RoomType{
				Type: models.Private,
			},
			RoomPassword: models.RoomPassword{
				Password: password,
			},
		}
	}

	db.Create(&room)
	
	db.Model(&room).Update("name", "Рум #" + strconv.Itoa(int(room.ID)))
	
	return c.JSON(http.StatusCreated, room)
}

func GetRooms(c echo.Context) error {
	var rooms []models.Room

    db.Preload("RoomPassword").Preload("RoomType").Find(&rooms)

    return c.JSON(http.StatusOK, rooms)
}


func EnterRoom(c echo.Context) error {
	var room models.Room
	db.Preload("RoomType").Preload("RoomPassword").Where("name = ?", c.FormValue("name")).First(&room)

	if room.RoomType.Type == "private" {
		password := c.FormValue("password")
		if password != room.RoomPassword.Password {
			return c.String(http.StatusBadRequest, "Invalid room password")
		}
		return c.JSON(http.StatusOK, room)
	}

	return c.JSON(http.StatusOK, room)
}

func FindRoomByName(c echo.Context) error {
	var rooms []models.Room
	db.Preload("RoomPassword").Preload("RoomType").Where("name LIKE ?", "%" + c.Param("name") + "%").Find(&rooms)

	return c.JSON(http.StatusOK, rooms)
}

func DeleteRoom(c echo.Context) error {
	var room models.Room
    db.Preload("RoomPassword").Preload("RoomType").Take(&room, c.Param("id"))

	if room.ID == 0 {
		return c.String(http.StatusNotFound, "no room found")
	}

    db.Delete(&room)

    return c.String(http.StatusOK, "room deleted")
}