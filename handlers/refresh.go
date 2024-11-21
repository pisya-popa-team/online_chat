package handlers

import (
	"net/http"
	"online_chat/enviroment"
	"online_chat/service"
	"online_chat/utils"

	"github.com/labstack/echo/v4"
)

func RefreshTokens(c echo.Context) error {
	token := utils.ExtractTokenFromHeaderString(c.Request().Header.Get("Authorization"))
    id := service.ExtractUsernameFromToken(token, enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET"))

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "0",
		"tokens": map[string]string{
			"access_token": service.NewAccessToken(id),
            "refresh_token": service.NewRefreshToken(id),
		},
	})

}
