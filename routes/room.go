package routes

import (
	"online_chat/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoomRoutes(access *echo.Group) {
	access.POST("/room", handlers.CreateRoom)
	access.GET("/rooms", handlers.GetRooms)
	access.POST("/room/enter", handlers.EnterRoom)
	access.DELETE("/room/:id", handlers.DeleteRoom)
	access.GET("/room/:name", handlers.FindRoomByName)
}