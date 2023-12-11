package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

const port = ":8089"

func main() {

	r := gin.Default()

	s, err := NewStore()
	if err != nil {
		log.Fatal(err)
	}

	r = setupRouter(r, s)

	err = r.Run(port)
	if err != nil {
		log.Fatal(err)
	}

}

func setupRouter(r *gin.Engine, s *Store) *gin.Engine {

	a := NewHandlerA(s)
	b := NewHandlerB(s)

	r.GET("/a", a.Run)
	r.GET("/b", b.Run)

	return r
}
