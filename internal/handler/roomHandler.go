package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DefaultRoomHandler(c *gin.Context) {

	c.JSON(http.StatusOK, h.Store.DefaultRooms)

}
