package router

import (
	"github.com/gin-gonic/gin"
	"mirrorb.fun/hello-artist/controller"
)

func InitArtistRouter(router *gin.Engine) {
	artistRouter := router.Group("/artist")
	{
		artistRouter.GET("/:name", controller.FindArtistByName)
	}
}
