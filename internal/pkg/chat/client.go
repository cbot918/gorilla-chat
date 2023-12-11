package chat

import (
	"gorilla-chat/internal/types"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections by disabling the CheckOrigin function
		return true
	},
}

type Client struct {
	ID    string
	Name  string
	Email string
	Conn  *websocket.Conn
}

func NewClient(u *types.User, w http.ResponseWriter, r *http.Request) *Client {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &Client{u.ID, u.Name, u.Email, conn}
}
