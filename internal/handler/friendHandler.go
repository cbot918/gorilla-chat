package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageType int

const (
	AddFriend MessageType = iota
)

type addFriendRequest struct {
	From string `json:"from"`
	Name string `json:"name"`
}

type toWSRequest struct {
	MType MessageType `json:"mtype"`
	From  string      `json:"from"`
}

func (h *Handler) AddFriendHandler(c *gin.Context) {

	var req addFriendRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := h.Dao.FindUserByName(req.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if !ok {
		err := fmt.Errorf("this user not exists")
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	fmt.Println(h.Store.Clients[req.Name])

	msgToTarget := &toWSRequest{
		MType: AddFriend,
		From:  req.From,
	}

	err = h.Store.Clients[req.Name].Conn.WriteJSON(msgToTarget)
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "已對 " + req.Name + " 送出申請"})

}
