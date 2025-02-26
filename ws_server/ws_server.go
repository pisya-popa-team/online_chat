package wsserver

import (
	"fmt"
	"net/http"
	// "sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
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

func ServeWs(c echo.Context) error {
    conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status": "5",
            "error": "failed to upgrade connection",
		})
	}

	defer conn.Close()

	client := newClient(conn)

	fmt.Println("New Client joined the hub!")
	fmt.Println(client)

	for {
		// Write
		err := conn.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := conn.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}