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

	routes.InitRegRoute(e)
	routes.InitAuthRoute(e)
	routes.InitRefreshRoute(refresh)
	routes.InitUserRoutes(access)
	routes.InitRoomRoutes(access)
	routes.InitRecoverRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}