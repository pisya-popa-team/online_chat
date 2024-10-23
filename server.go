package main

import (
	"online_chat/enviroment"
	"online_chat/handlers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	auth := e.Group("/auth")

	auth.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET")),
	}))

	e.POST("/user", handlers.CreateUser)

	auth.GET("/user/:id", handlers.GetUserByID)
	auth.GET("/users", handlers.GetAllUsers)
	auth.PUT("/user/:id", handlers.UpdateUser)
	auth.DELETE("/user/:id", handlers.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}