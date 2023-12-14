package handler

import (
	"fmt"
	"gorilla-chat/internal/util"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type RoomMessageResp struct {
	ID        int       `json:"room_id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *Handler) RoomHistoryMessageHandler(c *gin.Context) {

	fmt.Println("hihi")

	rommID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	messages, err := h.Dao.GetRoomMessages(rommID)
	if err != nil {
		fmt.Println("get room message failed")
		fmt.Println(err)
		return
	}

	results := []RoomMessageResp{}
	for _, m := range messages {
		results = append(results, RoomMessageResp{
			m.ID, m.Name, m.Content, m.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, results)

}

type ReceiveMessage struct {
	RoomID  int    `json:"room_id"`
	UserID  int    `json:"user_id"`
	EmailID string `json:"email"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (h *Handler) ReceiveRoomMessageHandler(c *gin.Context) {

	var receivedMessage ReceiveMessage
	err := c.BindJSON(&receivedMessage)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	util.PrintJSON(receivedMessage)

}
