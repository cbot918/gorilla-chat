package main

import (
	"gorilla-chat/pkg"
	"gorilla-chat/pkg/store"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const (
	port        = ":8088"
	webUrl      = "/"
	assetFolder = "ui/dist"
	wsUrl       = "/ws"
)

func main() {

	r := gin.Default()

	cfg, err := pkg.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := pkg.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	s, err := store.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	r = setupHTTPServer(r, db, s)

	err = r.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}

func setupHTTPServer(e *gin.Engine, db *sqlx.DB, store *store.Store) *gin.Engine {

	e.Use(myCorsPolicy())

	e = pkg.SetupWEB(e, webUrl, assetFolder)
	e = pkg.SetupAPIRouter(e, db, store)
	e = pkg.SetupSocketRouter(e, wsUrl, store)

	return e
}

func defaultCorsPolicy() gin.HandlerFunc {
	return cors.Default()
}

func myCorsPolicy() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization") // Add other headers here if needed
	config.AddAllowMethods("GET", "POST")   // Add other HTTP methods here if needed
	return cors.New(config)
}