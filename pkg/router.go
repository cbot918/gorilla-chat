package pkg

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupWEB(r *gin.Engine, urlPath string, assetPath string) *gin.Engine {

	r.Use(static.Serve(urlPath, static.LocalFile(assetPath, true)))

	return r
}

func SetupSocketRouter(r *gin.Engine, path string) *gin.Engine {

	r.GET(path, HandleWS)

	return r
}

func SetupAPIRouter(r *gin.Engine) *gin.Engine {

	api := r.Group("/api")
	{

		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "world"})
		})
	}

	r.GET("/ping", Ping)

	return r
}
