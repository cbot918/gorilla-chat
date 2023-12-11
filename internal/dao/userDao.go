package dao

import (
	"encoding/json"
	"fmt"
	"gorilla-chat/internal/types"
)

func (d *Dao) GetAllUserName() ([]types.User, error) {

	users := []types.User{}

	err := d.DB.Select(&users, "SELECT name FROM users")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return users, nil
}

func PrintJSON(v any) {
	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("json transform failed")
		return
	}
	fmt.Println(string(json))
}
