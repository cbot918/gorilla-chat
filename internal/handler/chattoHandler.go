package handler

import (
	"gorilla-chat/internal/util"
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

	util.PrintJSON(chattoQuery)

	c.JSON(http.StatusOK, chattoQuery)
}
