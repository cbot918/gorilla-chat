package pkg

import (
	"database/sql"
	"fmt"

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

func (d *Dao) CreateUser(name string, email string, password string) error {

	query := `INSERT INTO users (name, email, password) VALUES (?,?,?)`
	_, err := d.DB.Exec(query, name, email, password)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) GetUser(email string, password string) (User, error) {
	var user User
	err := d.DB.Get(&user, "SELECT * FROM users WHERE email=?", email)
	if err != nil {
		return User{}, err
	}
	if user.Password != password {
		return User{}, fmt.Errorf("signin failed")
	}
	return user, nil
}

func (d *Dao) userExists(email string) (bool, error) {
	var user User
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

func (d *Dao) GetUserByID(id string) (User, error) {
	var user User
	err := d.DB.Get(&user, "SELECT * FROM users WHERE id=?", id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
