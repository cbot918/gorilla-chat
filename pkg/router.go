package pkg

import (
	"gorilla-chat/pkg/handler"

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

	h := handler.NewHandler(db)

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

	// // create, delete, invite, join
	// room := e.Group("/room")
	// {

	// }

	// users: onlineUsers, offlineUsers
	users := e.Group("/user")
	{
		users.GET("/online", h.OnlineUser)
		users.GET("/all", h.AllUser)
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
