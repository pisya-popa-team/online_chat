package routes

import (
	wsserver "online_chat/ws_server"

	"github.com/labstack/echo/v4"
)

func InitWsRoute(e *echo.Echo){
	e.GET("/ws", wsserver.ServeWs)
}
