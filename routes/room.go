package routes

import (
	"online_chat/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoomRoutes(access *echo.Group) {
	access.POST("/rooms", handlers.CreateRoom)
	access.GET("/rooms", handlers.GetRooms)
	access.POST("/rooms/:id", handlers.EnterRoom)
	access.DELETE("/rooms/:id", handlers.DeleteRoom)
	access.GET("/rooms/:name", handlers.FindRoomByName)
}