package main

import (
	"online_chat/enviroment"
	"online_chat/routes"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	access := e.Group("/access")
	access.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET")),
	}))

	refresh := e.Group("/refresh")
	refresh.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(enviroment.GoDotEnvVariable("REFRESH_TOKEN_SECRET")),
	}))
	room := e.Group("/access/rooms")

	routes.InitRegRoute(e)
	routes.InitAuthRoute(e)
	routes.InitRecoverRoutes(e)
	routes.InitRefreshRoute(refresh)
	routes.InitUserRoutes(access)
	routes.InitRoomRoutes(access)
	routes.InitMessageRoutes(room)

	e.Logger.Fatal(e.Start(":1323"))
}