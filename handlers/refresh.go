package handlers

import (
	"net/http"
	"online_chat/enviroment"
	"online_chat/service"

	"github.com/labstack/echo/v4"
)

var (
	refresh_secret = enviroment.GoDotEnvVariable("REFRESH_TOKEN_SECRET")
)

func RefreshTokens(c echo.Context) error {
	refresh_token := c.FormValue("refresh_token")
	parsed_token, _ := service.ParseToken(refresh_token, refresh_secret)

	if !parsed_token.Valid {
		return c.String(http.StatusUnauthorized, "token is invalid")
	}
	
	username := service.ExtractUsernameFromToken(refresh_token, refresh_secret)

	return c.JSON(http.StatusCreated, map[string]string {
		"access_token": service.NewAccessToken(username),
		"refresh_token": service.NewRefreshToken(username),
	})

}
