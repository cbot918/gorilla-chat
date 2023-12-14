package handler

import (
	"fmt"
	"gorilla-chat/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type roomResponse struct {
	RoomID   int    `json:"room_id"`
	RoomName string `json:"room_name"`
}

func (h *Handler) DefaultRoomHandler(c *gin.Context) {

	var roomResponse []types.Room

	err := h.Dao.DB.Select(&roomResponse, "SELECT * FROM rooms")
	if err != nil {
		fmt.Println(err)
		// c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, roomResponse)
}

type EnterRoomRequest struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	RoomID int    `json:"room_id"`
}

func (h *Handler) EnterRoomHandler(c *gin.Context) {

	var req EnterRoomRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	h.Store.Clients[req.Name].CurrentRoom = req.RoomID

	c.JSON(http.StatusOK, gin.H{"msg": "enter room success"})
}
