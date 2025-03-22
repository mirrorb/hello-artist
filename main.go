package main

import (
	"mirrorb.fun/hello-artist/db"
	"mirrorb.fun/hello-artist/router"
)

func main() {
	db.InitDB()
	r := router.InitRouter()
	r.Run(":8080")
}
