package main

import (
	"online_chat/enviroment"
	"online_chat/routes"


	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()
	access := e.Group("/access")
	access.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET")),
	}))

	routes.InitRegRoute(e)
	routes.InitAuthRoute(e)
	routes.InitRefreshRoute(e)
	routes.InitUserRoutes(access)
	routes.InitRoomRoutes(access)

	

	e.Logger.Fatal(e.Start(":1323"))
}