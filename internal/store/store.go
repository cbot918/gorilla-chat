package store

import (
	"encoding/json"
	"fmt"

	"gorilla-chat/internal/pkg/chat"
	"gorilla-chat/internal/types"
	"gorilla-chat/internal/util"
)

type Store struct {
	Clients      map[string]*chat.Client
	OnlineUsers  []string
	DefaultRooms []types.DefaultRoom
}

func NewStore() (*Store, error) {

	// use default_room as group room
	jsonBytes, err := util.ReadJSON("default_room.json")
	if err != nil {
		fmt.Println("read default_room failed")
		return nil, err
	}

	rooms := []types.DefaultRoom{}
	err = json.Unmarshal(jsonBytes, &rooms)
	if err != nil {
		fmt.Println("unmarshal default_room failed")
		return nil, err
	}

	return &Store{
		Clients:      make(map[string]*chat.Client),
		OnlineUsers:  []string{},
		DefaultRooms: rooms,
	}, nil
}

func (s *Store) AddOnlineUsers(name string) {
	s.OnlineUsers = append(s.OnlineUsers, name)
}

func (s *Store) PrintOnlineUsers() {
	fmt.Println(s.OnlineUsers)
}
