package handler

import (
	"fmt"
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

type AddChattoParam struct {
	UserID   int `json:"user_id"`
	ChattoID int `json:"chatto_id"`
}

func (h *Handler) AddChattoHandler(c *gin.Context) {

	var addChattoParam AddChattoParam
	err := c.Bind(&addChattoParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = h.Dao.AddChatTo(addChattoParam.UserID, addChattoParam.ChattoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println("should success")

	c.JSON(http.StatusOK, gin.H{"msg": "add success"})

}
