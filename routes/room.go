package routes

import (
	"online_chat/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoomRoutes(access *echo.Group) {
	access.POST("/room", handlers.CreateRoom)
	access.GET("/rooms", handlers.GetRooms)
	// access.GET("/room/:name", handlers.FindRoomByName)
	access.GET("/room/:id", handlers.EnterRoom)
	access.DELETE("/room/:id", handlers.DeleteRoom)
}