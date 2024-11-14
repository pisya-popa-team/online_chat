package handlers

import (
	"net/http"
	"online_chat/enviroment"
	"online_chat/service"
	"online_chat/utils"

	"github.com/labstack/echo/v4"
)

var (
	refresh_secret = enviroment.GoDotEnvVariable("REFRESH_TOKEN_SECRET")
)

func RefreshTokens(c echo.Context) error {
	refresh_token := c.FormValue("refresh_token")
	parsed_token, _ := service.ParseToken(refresh_token, refresh_secret)

	if !parsed_token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"status": "1",
			"error": "token is invalid",
		})
	}
	
	id := service.ExtractUsernameFromToken(refresh_token, refresh_secret)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": 0,
		"tokens": map[string]string{
			"access_token": service.NewAccessToken(utils.StringToUint(id)),
            "refresh_token": service.NewRefreshToken(utils.StringToUint(id)),
		},
	})

}
