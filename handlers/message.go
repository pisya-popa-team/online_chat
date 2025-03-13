package handlers

import (
	"net/http"
	"online_chat/models"
	"online_chat/utils"

	"github.com/labstack/echo/v4"
)

func GetMessages(c echo.Context) error {
	var messages []models.Message
	limit, offset := utils.StringToInt(c.QueryParam("limit")), utils.StringToInt(c.QueryParam("offset"))

	db.Where("room_id = ?", c.Param("id")).Limit(limit).Offset(offset).Find(&messages)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "0",
		"messages": messages,
	})
}