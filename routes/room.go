package routes

import (
	"online_chat/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoomRoutes(access *echo.Group) {
	access.POST("/room", handlers.CreateRoom)
	access.GET("/rooms", handlers.GetRooms)
	access.DELETE("/room/:id", handlers.DeleteRoom)
}