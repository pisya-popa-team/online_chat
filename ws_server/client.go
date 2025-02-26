package wsserver

import (
	"online_chat/models"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	Rooms map[uint]*models.Room
}

func newClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
	}
}