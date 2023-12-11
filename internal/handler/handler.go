package handler

import (
	// "gorilla-chat/pkg"
	"gorilla-chat/internal/dao"
	"gorilla-chat/internal/store"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Dao   *dao.Dao
	Store *store.Store
}

func NewHandler(db *sqlx.DB, store *store.Store) *Handler {

	return &Handler{
		Dao:   dao.NewDao(db),
		Store: store,
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
