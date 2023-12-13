package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RoomMessageResp struct {
	ID        int       `json:"room_id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *Handler) RoomMessageHandler(c *gin.Context) {

	fmt.Println("hihi")

	rommID := 1

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
