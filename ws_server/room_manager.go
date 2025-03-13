package wsserver

import (
	"errors"
	"online_chat/models"
)

type RoomManager struct {
	Rooms   map[uint]*Room
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		Rooms: make(map[uint]*Room),
	}
}

func (rm *RoomManager) RemoveRoom (id uint) {
	delete(rm.Rooms, id)
}

func (rm *RoomManager) RemoveClientFromRoom(client *Client, room_id uint) {
	for i, this_client := range rm.Rooms[room_id].Clients {
		if this_client.UserID == client.UserID {
			rm.Rooms[room_id].Clients = append(rm.Rooms[room_id].Clients[:i], rm.Rooms[room_id].Clients[i+1:]...)
			break
		}
	}

	if len(rm.Rooms[room_id].Clients) == 0 {
		close(rm.Rooms[room_id].Messages)
		rm.RemoveRoom(room_id)
	}
}

func (rm *RoomManager) AddClientToRoom(client *Client, room_id uint) error {
	var this_room models.Room
	db.Where("id = ?", room_id).Find(&this_room)

	if this_room.ID == 0 {
        return errors.New("Room not found")
    }

	room, exists := rm.Rooms[room_id]
	if !exists {
		room = NewRoom(room_id)
		rm.Rooms[room_id] = room

		room.start_once.Do(func() {
			go room.HandleMessages(rm)
		})
	}

	for _, this_client := range room.Clients {
		if this_client.UserID == client.UserID {
			return nil
		}
	}

	room.Clients = append(room.Clients, client)
	
	go client.ReadMessages(room, rm)

	return nil
}