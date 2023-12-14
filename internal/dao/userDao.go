package dao

import (
	"gorilla-chat/internal/types"
)

func (d *Dao) GetAllUserName() ([]types.User, error) {

	users := []types.User{}

	err := d.DB.Select(&users, "SELECT name FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (d *Dao) GetUserNameByID(userID int) (string, error) {
	user := types.User{}

	err := d.DB.Get(&user, "SELECT name FROM users where user_id=?", userID)
	if err != nil {
		return "", err
	}

	return user.Name, nil
}
