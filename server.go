package main

import (
	"online_chat/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.InitRegRoute(e)
	routes.InitAuthRoute(e)
	routes.InitUserRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}