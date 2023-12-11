package pkg

import (
	"fmt"
	"gorilla-chat/pkg/dao"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Dao *dao.Dao
}

func NewHandler(db *sqlx.DB) *Handler {

	return &Handler{
		Dao: dao.NewDao(db),
	}
}

type addFriendRequest struct {
	TargetEmail string `json:"targetEmail"`
}

func (h *Handler) AddFriendHandler(c *gin.Context) {

	fmt.Println("in add friend")
}

// other handler
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "world"})
}
