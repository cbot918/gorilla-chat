package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type onlineUserRequest struct {
	TargetEmail string `json:"targetEmail"`
}

func (h *Handler) OnlineUser(c *gin.Context) {

	fmt.Println("in add friend")
}

type allUserResponse struct {
	Count int      `json:"count"`
	Names []string `json:"names"`
}

func (h *Handler) AllUser(c *gin.Context) {
	users, err := h.Dao.GetAllUserName()
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	names := []string{}

	for _, user := range users {
		names = append(names, user.Name)
	}

	c.JSON(http.StatusOK, &allUserResponse{
		Names: names,
		Count: len(names),
	})
}
