package routes

import (
	"online_chat/handlers"
	"github.com/labstack/echo/v4"
)

func InitUserRoutes(access *echo.Group) {
	access.GET("/users/me", handlers.GetInfoAboutMe)
	access.GET("/user/:id", handlers.GetUserByID)
	access.GET("/users", handlers.GetAllUsers)
	access.PUT("/user/:id", handlers.UpdateUser)
	access.DELETE("/user/:id", handlers.DeleteUser)
}