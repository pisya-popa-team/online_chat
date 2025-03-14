package wsserver

import (
	"sync"
)

type Room struct {
	RoomID     uint
	Clients    []*Client
	Messages   chan ClientMessage
	start_once sync.Once
}

func NewRoom(room_id uint) *Room {
	return &Room{
		RoomID:  room_id,
		Clients: []*Client{},
		Messages: make(chan ClientMessage),
	}
}

type ClientMessage struct {
	MessageType  string `json:"message_type"`
	Content      string	`json:"content"`
	SentAt       string	`json:"sent_at"`
	RoomID       uint	`json:"room_id"` 
	UserID       uint	`json:"user_id"`
	Username     string	`json:"username"`
}

func (r *Room) HandleMessages(rm *RoomManager) {
	for client_message := range r.Messages {
		if len(r.Clients) == 0 {
			close(r.Messages)
			rm.RemoveRoom(r.RoomID)
			return
		}

		for _, client := range r.Clients {
			err := client.Conn.WriteJSON(client_message)
			
			if err != nil {
				client.Conn.Close()
				rm.RemoveClientFromRoom(client, r.RoomID)
			}
        }
	}
}
