package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Gin!")
	})
	pprof.Register(r)
	r.Run(":8080")
}
