package routes

import (
	"online_chat/handlers"
	"github.com/labstack/echo/v4"
)

func InitUserRoutes(access *echo.Group) {
	access.GET("/users/me", handlers.GetInfoAboutMe)
	access.GET("/users", handlers.GetAllUsers)
	access.PATCH("/user/update", handlers.UpdateUser)
}