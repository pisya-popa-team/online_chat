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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		fmt.Println("Request Origin:", origin) // Логирование
		allowedOrigins := map[string]bool{
			"http://localhost:4200": true,
			"https://tt-chat.danyatochka.ru": true,
			"https://api-tt-chat.danyatochka.ru": true,
		}
		return allowedOrigins[origin]
	},
}


var room_manager = NewRoomManager()
var secret = enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET")

func ServeWs(c echo.Context) error {
	token := c.QueryParam("token")
	room_id := uint(utils.StringToInt(c.Param("id")))
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
		return err
	}

	client := &Client{
		UserID: uint(utils.StringToInt(service.ExtractUsernameFromToken(token, secret))),
		Conn:   conn,
	}


	err = room_manager.AddClientToRoom(client, room_id)
	if err != nil {
		error_message := map[string]interface{}{
			"message": err.Error(),
		}
	
		_ = conn.WriteJSON(error_message)
		conn.Close()
		return nil
	}

	return nil
}

