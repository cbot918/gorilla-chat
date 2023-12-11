package dao

import (
	"github.com/jmoiron/sqlx"
)

type Dao struct {
	DB *sqlx.DB
}

func NewDao(db *sqlx.DB) *Dao {
	return &Dao{
		DB: db,
	}
}
