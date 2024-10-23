package main

import (
	"online_chat/enviroment"
	"online_chat/handlers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	access := e.Group("/access")

	access.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET")),
	}))

	e.POST("/reg", handlers.CreateUser)
	e.POST("/auth", handlers.Authorisation)

	access.GET("/user/:id", handlers.GetUserByID)
	access.GET("/users", handlers.GetAllUsers)
	access.PUT("/user/:id", handlers.UpdateUser)
	access.DELETE("/user/:id", handlers.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}