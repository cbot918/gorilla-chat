package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type onlineUserRequest struct {
	TargetEmail string `json:"targetEmail"`
}

type onlineUserResponse struct {
	Count int      `json:"count"`
	Users []string `json:"users"`
}

func (h *Handler) OnlineUser(c *gin.Context) {

	userResponse := &onlineUserResponse{}

	for _, v := range h.Store.Clients {
		userResponse.Users = append(userResponse.Users, v.Name)
	}
	userResponse.Count = len(userResponse.Users)

	c.JSON(http.StatusOK, userResponse)
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
