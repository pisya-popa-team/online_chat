package wsserver

import (
	"fmt"
	"net/http"
	"online_chat/enviroment"
	"online_chat/service"
	"online_chat/utils"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var allowedOrigins = map[string]bool{
	"http://localhost:4200": true,
	"https://tt-chat.danyatochka.ru": true,
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println(r.Header.Get("Origin"))
		_, ok := allowedOrigins[r.Header.Get("Origin")]
		return ok
	},
}


var room_manager = NewRoomManager()
var secret = enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET")

func ServeWs(c echo.Context) error {
	fmt.Println("ServeWs called!")
	token := c.QueryParam("token")
	room_id := uint(utils.StringToInt(c.Param("id")))
	fmt.Println("Token:", token)
    fmt.Println("Room ID:", room_id)
	err := service.ValidateAccessToken(token, secret)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status": 1,
			"message": err.Error(),
		})
	}

    conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
		fmt.Println("WebSocket Upgrade Error:", err)
		return err
	}

	fmt.Println("WebSocket connection established")

	client := &Client{
		UserID: uint(utils.StringToInt(service.ExtractUsernameFromToken(token, secret))),
		Conn:   conn,
	}


	err = room_manager.AddClientToRoom(client, room_id)
	if err != nil {
		fmt.Println("Error adding client to room:", err)
		error_message := map[string]interface{}{
			"message": err.Error(),
		}
	
		_ = conn.WriteJSON(error_message)
		conn.Close()
		return nil
	}

	return nil
}

