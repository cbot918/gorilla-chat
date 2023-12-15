package handler

import (
	"gorilla-chat/internal/util"
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

type onlineUserResp struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

func (h *Handler) OnlineUser(c *gin.Context) {

	onlineUserResps := []onlineUserResp{}

	for _, v := range h.Store.Clients {
		onlineUserResps = append(onlineUserResps, onlineUserResp{
			UserID: v.ID,
			Name:   v.Name,
		})
	}

	util.PrintJSON(onlineUserResps)

	c.JSON(http.StatusOK, onlineUserResps)
}

type allUserResponse struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

func (h *Handler) AllUser(c *gin.Context) {
	users, err := h.Dao.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var userResp []allUserResponse
	for _, user := range users {
		userResp = append(userResp, allUserResponse{
			UserID: user.ID,
			Name:   user.Name,
		})
	}

	c.JSON(http.StatusOK, userResp)
}
