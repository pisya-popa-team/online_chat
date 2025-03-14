package handlers

import (
	"net/http"
	"online_chat/utils"

	"github.com/labstack/echo/v4"
)

type MessageResponse struct {
	MessageType string `json:"message_type"`
	Content     string `json:"content"`
	SentAt      string `json:"sent_at"`
	RoomID      uint   `json:"room_id"`
	UserID      uint   `json:"user_id"`
	Username    string `json:"username"`
}

func GetMessages(c echo.Context) error {
	var messages []MessageResponse
	limit, offset := utils.StringToInt(c.QueryParam("limit")), utils.StringToInt(c.QueryParam("offset"))

	db.Table("messages").
    Select("messages.message_type, messages.content, messages.sent_at, messages.room_id, messages.user_id, users.username").
    Joins("LEFT JOIN users ON users.id = messages.user_id").Where("messages.room_id = ?", c.Param("id")).Limit(limit).Offset(offset).Scan(&messages)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "0",
		"messages": messages,
	})
}