package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HandlerA struct {
	S *Store
}

func NewHandlerA(s *Store) *HandlerA {
	return &HandlerA{s}
}

func (h *HandlerA) Run(c *gin.Context) {
	fmt.Println("in A")
	h.S.AddOnlineUsers("hihi")
	h.S.PrintOnlineUsers()
}

type HandlerB struct {
	S *Store
}

func NewHandlerB(s *Store) *HandlerB {
	return &HandlerB{s}
}

func (h *HandlerB) Run(c *gin.Context) {
	fmt.Println("in B")
	h.S.AddOnlineUsers("hoho")
	h.S.PrintOnlineUsers()
}
