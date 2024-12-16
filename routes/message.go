package routes

import (
	"online_chat/handlers"

	"github.com/labstack/echo/v4"
)

func InitMessageRoutes(room *echo.Group){
	room.GET("/:id/messages", handlers.GetMessages)
}