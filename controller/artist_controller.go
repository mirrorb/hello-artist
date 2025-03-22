package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mirrorb.fun/hello-artist/service"
)

func FindArtistByName(c *gin.Context) {
	name := c.Param("name")
	artists, err := service.FindArtistByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, artists)
}
