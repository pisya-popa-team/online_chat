package routes

import (
	"online_chat/handlers"

	"github.com/labstack/echo/v4"
)

func InitRefreshRoute(e *echo.Echo) {
    e.POST("/refresh", handlers.RefreshTokens)
}