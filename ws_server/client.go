package wsserver

import (
	"online_chat/models"
	"online_chat/utils"

	"github.com/gorilla/websocket"
)

type Client struct {
	UserID    uint
	Conn      *websocket.Conn
}

type ClientMessage struct {
	MessageType  string      `json:"message_type"`
	Content      string      `json:"content"`
	SentAt       string	     `json:"sent_at"`
	UserName     string      `json:"username"`
	RoomID       uint        `json:"room_id"`
}

func (c *Client) ReadMessages(room *Room, room_manager *RoomManager) {
	defer func() {
		room_manager.RemoveClientFromRoom(c, room.RoomID)
		c.Conn.Close()
	}()

	for {
		var client_message ClientMessage
		err := c.Conn.ReadJSON(&client_message)
		if err != nil {
			break
		}

		message := models.Message{
			Content: client_message.Content,
			SentAt: utils.FormatStringToDate(client_message.SentAt),
			RoomID: client_message.RoomID,
			UserID: uint(utils.StringToInt(c.FindUserName())),
		}

		db.Create(&message)

		room.Messages <- client_message
	}
}

func (c *Client) FindUserName () string {
	var user models.User
	db.Where("id = ?", c.UserID).Find(&user)
	return user.Username
}