package pkg

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func HandleWS(c *gin.Context) {

	user := &User{
		ID:    "123",
		Name:  "yale",
		Email: "yale918@gmail.com",
	}

	client := NewClient(user, c.Writer, c.Request)

	mType, msg, err := client.Conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	err = client.Conn.WriteMessage(mType, msg)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(msg))
}
