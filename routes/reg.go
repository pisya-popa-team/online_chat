package routes

import (
	"online_chat/handlers"

	"github.com/labstack/echo/v4"
)


func InitRegRoute(e *echo.Echo) {
	e.POST("/reg", handlers.CreateUser)
}