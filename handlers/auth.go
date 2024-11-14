package handlers

import (
	"net/http"
	"online_chat/models"
	"online_chat/password_hashing"
	"online_chat/service"

	"github.com/labstack/echo/v4"
)

func Authorisation(c echo.Context) error {
	email, password := c.FormValue("email"), c.FormValue("password")

	var user models.User

	db.Preload("Password").Where("email = ?", email).Find(&user)

	if user.ID == 0 || !password_hashing.DoPasswordsMatch(user.Password.Hash, password){
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"status": "1",
			"error": "invalid user info",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 0,
		"tokens": map[string]string{
			"access_token": service.NewAccessToken(user.ID),
            "refresh_token": service.NewRefreshToken(user.ID),
		},
	})
}