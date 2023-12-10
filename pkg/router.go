package pkg

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupWEB(r *gin.Engine, urlPath string, assetPath string) *gin.Engine {

	r.Use(static.Serve(urlPath, static.LocalFile(assetPath, true)))

	return r
}

func SetupSocketRouter(r *gin.Engine, path string, db *sqlx.DB) *gin.Engine {

	// socket := r.Group("/ws")
	// socket.Use(RequireLogin(db))
	r.GET(path, HandleWS)

	return r
}

func SetupAPIRouter(e *gin.Engine, db *sqlx.DB) *gin.Engine {

	h := NewHandler(db)

	auth := e.Group("/auth")
	{

		auth.POST("/signup", h.SignupHandler)
		auth.POST("/signin", h.SigninHandler)
		auth.POST("/authbeforews", h.AuthBeforeWSHandler)
	}

	// 1v1, 1vn, broadcast
	message := e.Group("/message")
	{
		message.POST("")
	}

	// create, delete, invite, join
	room := e.Group("/room")
	{

	}

	e.GET("/hello", Hello)
	e.GET("/ping", Ping)

	return e
}
