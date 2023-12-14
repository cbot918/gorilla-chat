package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DefaultRoomHandler(c *gin.Context) {

	c.JSON(http.StatusOK, h.Store.DefaultRooms)

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
