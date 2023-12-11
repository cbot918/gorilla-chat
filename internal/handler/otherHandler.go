package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// other handler
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *Handler) Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "world"})
}
