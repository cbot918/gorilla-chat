package handler

import (
	"gorilla-chat/pkg/dao"

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

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
