package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type addFriendRequest struct {
	TargetEmail string `json:"targetEmail"`
}

func (h *Handler) AddFriendHandler(c *gin.Context) {

	fmt.Println("in add friend")
}
