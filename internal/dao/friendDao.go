package dao

import (
	"database/sql"
	"gorilla-chat/internal/types"
)

func (d *Dao) FindUserByName(name string) (bool, error) {
	var user types.User
	err := d.DB.Get(&user, "SELECT * FROM users WHERE name=?", name)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return true, err
		}
	}
	return true, nil
}
