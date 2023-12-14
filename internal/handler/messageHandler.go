package handler

import (
	"fmt"
	"gorilla-chat/internal/util"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// server received client message post
type ReceiveMessage struct {
	RoomID   int    `json:"room_id"`
	UserID   int    `json:"user_id"`
	EmailID  string `json:"email"`
	Name     string `json:"name"`
	Message  string `json:"message"`
	ToUserID int    `json:"to_user_id"`
}

func (h *Handler) ReceiveRoomMessageHandler(c *gin.Context) {

	var req ReceiveMessage
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = h.Dao.AddMessageHistory(req.RoomID, req.UserID, req.ToUserID, req.Message)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	util.PrintJSON(req)

	// 1 v 1 聊天
	if req.RoomID == 0 {
		toUserName, err := h.Dao.GetUserNameByID(req.ToUserID)

		if err != nil {
			fmt.Println("heree")
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		h.Store.Clients[req.Name].Conn.WriteJSON(req)

		_, ok := h.Store.Clients[toUserName]
		if ok {
			h.Store.Clients[toUserName].Conn.WriteJSON(req)
		}

	} else { // 群聊
		for _, c := range h.Store.Clients {
			if c.CurrentRoom == req.RoomID {
				c.Conn.WriteJSON(req)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"reply": "http message received"})
}

// room history message
type RoomMessageResp struct {
	ID        int       `json:"room_id"`
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *Handler) RoomHistoryMessageHandler(c *gin.Context) {

	rommID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	messages, err := h.Dao.GetRoomHistoryMessages(rommID)
	if err != nil {
		fmt.Println("get room message failed")
		fmt.Println(err)
		return
	}

	results := []RoomMessageResp{}
	for _, m := range messages {
		results = append(results, RoomMessageResp{
			m.ID, m.UserID, m.Name, m.Content, m.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, results)

}

// user history message
type UserHistoryMessageParam struct {
	UserID   int `json:"user_id"`
	ToUserID int `json:"to_user_id"`
}

func (h *Handler) UserHistoryMessageHandler(c *gin.Context) {

	var req UserHistoryMessageParam

	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userHistory, err := h.Dao.GetUserHistoryMessages(req.UserID, req.ToUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if userHistory == nil {
		c.JSON(http.StatusOK, []string{})
		return
	}

	c.JSON(http.StatusOK, userHistory)

}
