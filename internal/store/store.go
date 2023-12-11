package store

import (
	"fmt"
	"gorilla-chat/internal/pkg/chat"
)

type Store struct {
	Clients     map[string]*chat.Client
	OnlineUsers []string
}

func NewStore() (*Store, error) {
	return &Store{
		Clients:     make(map[string]*chat.Client),
		OnlineUsers: []string{},
	}, nil
}

func (s *Store) AddOnlineUsers(name string) {
	s.OnlineUsers = append(s.OnlineUsers, name)
}

func (s *Store) PrintOnlineUsers() {
	fmt.Println(s.OnlineUsers)
}
