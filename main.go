package main

import (
	// "net/http"
	"online_chat/enviroment"
	"online_chat/routes"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	e := echo.New()

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://api-tt-chat.danyatochka.ru", "http://localhost:4200"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	// 	AllowCredentials: true,
	// }))

	e.Use(middleware.CORS())

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