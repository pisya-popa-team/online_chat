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
	id := service.ExtractUsernameFromToken(token, enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET"))
	
	var (
		user models.User
		room models.Room
	)

	db.Where("id = ?", id).Find(&user)

	password := c.FormValue("password")
	if password == "" {
		room = models.Room{
			UserID: user.ID,
			RoomType: models.Public,
		}
	} else {
		room = models.Room{
			UserID: user.ID,
			RoomType: models.Private,
			RoomPassword: models.RoomPassword{
				Password: password,
			},
		}
	}

	db.Create(&room)
	
	db.Model(&room).Update("name", "Рум #" + strconv.Itoa(int(room.ID)))
	
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "0",
		"room": room,
	})
}

func GetRooms(c echo.Context) error {
	var rooms []models.Room

    db.Find(&rooms)

    return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "0",
		"rooms": rooms,
	})
}


func EnterRoom(c echo.Context) error {
	var (
		room models.Room
		user models.User
	)
	db.Preload("RoomPassword").Preload("Users").Where("id = ?", c.Param("id")).First(&room)

	if room.RoomType == "private" {
		password := c.FormValue("password")
		if password != room.RoomPassword.Password {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"status": "1",
				"message": "invalid room password",
			})
		}
	}

	token := utils.ExtractTokenFromHeaderString(c.Request().Header.Get("Authorization"))
	id := service.ExtractUsernameFromToken(token, enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET"))

	db.Where("id = ?", id).Find(&user)
	db.Model(&room).Association("Users").Append(&user)
	count := db.Model(&room).Association("Users").Count()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "0",
		"room": room,
		"count": count,
	})
}

func FindRoomByName(c echo.Context) error {
	var rooms []models.Room
	db.Where("name LIKE ?", "%" + c.Param("name") + "%").Find(&rooms)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "0",
		"rooms": rooms,
	})
}

func DeleteRoom(c echo.Context) error {
	token := utils.ExtractTokenFromHeaderString(c.Request().Header.Get("Authorization"))
	user_id := service.ExtractUsernameFromToken(token, enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET"))

	var room models.Room
	
    db.Where("user_id = ? AND id = ?", user_id, c.Param("id")).Find(&room)

	if room.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
            "status": "3",
            "error": "no room found",
        })
	}

    db.Select("RoomPassword").Delete(&room)

    return c.JSON(http.StatusOK, map[string]string{
		"status": "0",
		"message": "room deleted",
	})
}