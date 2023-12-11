package dao

import (
	"database/sql"
	"fmt"
	"gorilla-chat/pkg/types"
)

func (d *Dao) CreateUser(name string, email string, password string) error {

	query := `INSERT INTO users (name, email, password) VALUES (?,?,?)`
	_, err := d.DB.Exec(query, name, email, password)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) GetUser(email string, password string) (types.User, error) {
	var user types.User
	err := d.DB.Get(&user, "SELECT * FROM users WHERE email=?", email)
	if err != nil {
		return types.User{}, err
	}
	if user.Password != password {
		return types.User{}, fmt.Errorf("signin failed")
	}
	return user, nil
}

func (d *Dao) EmailExists(email string) (bool, error) {
	var user types.User
	err := d.DB.Get(&user, "SELECT * FROM users WHERE email=?", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (d *Dao) NameExists(name string) (bool, error) {
	var user types.User
	err := d.DB.Get(&user, "SELECT * FROM users WHERE name=?", name)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return true, err
		}
	}
	return true, fmt.Errorf("name exists")
}

func (d *Dao) GetUserByID(id string) (types.User, error) {
	var user types.User
	err := d.DB.Get(&user, "SELECT * FROM users WHERE id=?", id)
	if err != nil {
		return types.User{}, err
	}

	return user, nil
}
