package internal

import (
	"fmt"
	"gorilla-chat/internal/pkg/chat"
	"gorilla-chat/internal/store"
	"gorilla-chat/internal/types"
	"log"

	"github.com/gin-gonic/gin"
)

type WSHandler struct {
	Store *store.Store
}

func NewWSHandler(s *store.Store) *WSHandler {
	return &WSHandler{
		Store: s,
	}
}

func (ws *WSHandler) HandleWS(c *gin.Context) {

	user := &types.User{
		ID:    c.Query("id"),
		Name:  c.Query("name"),
		Email: c.Query("email"),
	}

	client := chat.NewClient(user, c.Writer, c.Request)

	ws.Store.Clients[user.Name] = client

	// PrintJSON(ws.Store.Clients)

	go ReadWorker(client)

	// mType, msg, err := client.Conn.ReadMessage()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// err = client.Conn.WriteMessage(mType, msg)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

}

func ReadWorker(client *chat.Client) {
	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(msg))
	}
}

func WriteWorker() {

}
