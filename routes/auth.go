package routes

import (
	"online_chat/handlers"

	"github.com/labstack/echo/v4"
)

func InitAuthRoute(e *echo.Echo) {
	e.POST("/auth", handlers.Authorisation)
}