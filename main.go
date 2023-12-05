package main

import (
	"gorilla-chat/pkg"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	port        = ":8088"
	webUrl      = "/"
	assetFolder = "ui/dist"
	wsUrl       = "/ws"
)

var corsPolicy = cors.Default()

func main() {

	r := gin.Default()

	r = setups(r)

	err := r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func setups(r *gin.Engine) *gin.Engine {

	r = pkg.SetupWEB(r, webUrl, assetFolder)
	r = pkg.SetupAPIRouter(r)
	r = pkg.SetupSocketRouter(r, wsUrl)

	r.Use(corsPolicy)

	return r
}
