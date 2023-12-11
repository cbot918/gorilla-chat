package pkg

import (
	"fmt"
	"gorilla-chat/pkg/types"
	"log"

	"github.com/gin-gonic/gin"
)

func HandleWS(c *gin.Context) {

	user := &types.User{
		ID:    c.Query("id"),
		Name:  c.Query("name"),
		Email: c.Query("email"),
	}

	client := NewClient(user, c.Writer, c.Request)

	PrintJSON(client)

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

func ReadWorker(client *Client) {
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
