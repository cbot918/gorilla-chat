package main

import "fmt"

type Store struct {
	OnlineUsers []string
}

func NewStore() (*Store, error) {
	return &Store{}, nil
}

func (s *Store) AddOnlineUsers(name string) {
	s.OnlineUsers = append(s.OnlineUsers, name)
}

func (s *Store) PrintOnlineUsers() {
	fmt.Println(s.OnlineUsers)
}
