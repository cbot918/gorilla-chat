package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UserChannelsHandler(c *gin.Context) {

	err := h.Dao.GetUserChannels(1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
