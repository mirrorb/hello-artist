package router

import (
	"github.com/gin-gonic/gin"
	"mirrorb.fun/hello-artist/controller"
)

func InitArtistRouter(router *gin.Engine) {
	router.GET("/artist", controller.FindArtistByName)
}
