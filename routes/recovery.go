package routes

import (
	"online_chat/handlers"

	"github.com/labstack/echo/v4"
)

func InitRecoverRoutes(e *echo.Echo) {
	e.POST("/recovery", handlers.GenerateRecoveryToken)
	e.GET("/recovery/:token", handlers.GetRecoveryToken)
	e.PATCH("/recovery", handlers.UseRecoveryToken)
}