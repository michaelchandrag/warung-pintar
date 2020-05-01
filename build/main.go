package main

import (
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main () {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}