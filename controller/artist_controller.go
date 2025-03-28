package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mirrorb.fun/hello-artist/service"
)

func FindArtistByName(c *gin.Context) {
	name := c.Query("name")
	q_type := c.Query("type")
	var results any
	var err error
	if q_type == "all" {
		results, err = service.FindResultByName(name)
	} else if q_type == "artist" {
		results, err = service.FindArtistByName(name)
	} else if q_type == "character" {
		results, err = service.FindCharacterByName(name)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query type",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, results)
}
