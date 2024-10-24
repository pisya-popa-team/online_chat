package handlers

import (
	"fmt"
	"net/http"
	"online_chat/models"
	"online_chat/password_hashing"
	"online_chat/service"

	"github.com/labstack/echo/v4"
)

func Authorisation(c echo.Context) error {
	username, password := c.FormValue("username"), c.FormValue("password")

	var user models.User

	db.Preload("Password").Where("username = ?", username).Find(&user)

	if user.ID == 0 {
		return c.String(http.StatusUnauthorized, "invalid username")
	}

	fmt.Println(password_hashing.DoPasswordsMatch(user.Password.Hash, password))

	if !password_hashing.DoPasswordsMatch(user.Password.Hash, password){
		return c.String(http.StatusUnauthorized, "invalid password")
	}

	access := service.NewAccessToken(username)
    refresh := service.NewRefreshToken()

	return c.JSON(http.StatusOK, map[string]string{
  		"message": "welcome " + user.Username + "!",
        "access_token": access,
        "refresh_token": refresh,
    })
}