package internal

import (
	"gorilla-chat/internal/handler"
	"gorilla-chat/internal/store"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupWEB(r *gin.Engine, urlPath string, assetPath string) *gin.Engine {

	r.Use(static.Serve(urlPath, static.LocalFile(assetPath, true)))

	return r
}

func SetupSocketRouter(r *gin.Engine, path string, store *store.Store) *gin.Engine {

	// socket := r.Group("/ws")
	// socket.Use(RequireLogin(db))
	h := NewWSHandler(store)
	r.GET(path, h.HandleWS)

	return r
}

func SetupAPIRouter(e *gin.Engine, db *sqlx.DB, store *store.Store) *gin.Engine {

	h := handler.NewHandler(db, store)

	auth := e.Group("/auth")
	{
		auth.POST("/signup", h.SignupHandler)
		auth.POST("/signin", h.SigninHandler)
		auth.POST("/authbeforews", h.AuthBeforeWSHandler)
	}

	// users: onlineUsers, offlineUsers
	users := e.Group("/user")
	{
		users.GET("/online", h.OnlineUser)
		users.GET("/all", h.AllUser)
	}

	channel := e.Group("/channel")
	{
		channel.GET("/", h.UserChannelsHandler)
	}

	// 1v1, 1vn, broadcast
	message := e.Group("/message")
	{
		message.POST("")
	}

	// friend: add
	friend := e.Group("/friend")
	{
		friend.POST("/add", h.AddFriendHandler)
	}

	e.GET("/hello", h.Hello)
	e.GET("/ping", h.Ping)

	return e
}
