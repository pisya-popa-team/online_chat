package wsserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		allowedOrigins := map[string]bool{
			"http://localhost:4200": true,
			"https://tt-chat.danyatochka.ru/": true,
		}
		return allowedOrigins[r.Header.Get("Origin")]
	},
}

// type WsServer struct {
// 	clients    map[*Client]bool
// 	register   chan *Client
// 	unregister chan *Client
// 	mu 		   sync.Mutex
// }

// func NewWsServer() *WsServer {
//     return &WsServer{
// 		clients:    make(map[*Client]bool),
// 		register:   make(chan *Client),
// 		unregister: make(chan *Client),
// 	}
// }

type Message struct {
	Content string `json:"content"`
}

func ServeWs(c echo.Context) error {
    conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	defer conn.Close()

	for {
		// Read
		_, msg, err := conn.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
			return err
		}

		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			c.Logger().Error("Ошибка разбора JSON:", err)
			continue
		}

		fmt.Printf("Получено сообщение: %s\n", message.Content)

		response := Message{Content: "Принято: " + message.Content}
		respJSON, err := json.Marshal(response)
		if err != nil {
			c.Logger().Error("Ошибка кодирования JSON:", err)
			continue
		}

		err = conn.WriteMessage(websocket.TextMessage, respJSON)
		if err != nil {
			c.Logger().Error(err)
			return err
		}
	}
}