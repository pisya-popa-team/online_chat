package wsserver

import (
	"online_chat/models"

	"github.com/gorilla/websocket"
)

type Client struct {
	UserID    uint
	Conn      *websocket.Conn
}

func GetUser(user_id uint) models.User {
	var user models.User
	db.Where("id =?", user_id).Find(&user)
	return user
}

func (c *Client) ReadMessages(room *Room, room_manager *RoomManager) {
	defer func() {
		room_manager.RemoveClientFromRoom(c, room.RoomID)
		c.Conn.Close()
	}()

	for {
		var this_message ClientMessage
		err := c.Conn.ReadJSON(&this_message)
		if err != nil {
			break
		}

		message := models.Message{
			MessageType: models.UserM,
			Content: this_message.Content,
			SentAt:  this_message.SentAt,
			RoomID:  this_message.RoomID,
			UserID:  this_message.UserID,
			User:    GetUser(this_message.UserID),
		}

		db.Create(&message)

		room.Messages <- this_message
	}
}