package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChattoParam struct {
	UserID int `json:"user_id"`
}

func (h *Handler) ChattoHandler(c *gin.Context) {

	var chatto ChattoParam

	err := c.Bind(&chatto)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	chattoQuery, err := h.Dao.GetChatTo(chatto.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if chattoQuery == nil {
		c.JSON(http.StatusOK, []string{})
		return
	}
	c.JSON(http.StatusOK, chattoQuery)
}
