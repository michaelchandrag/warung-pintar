package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/michaelchandrag/warung-pintar/module/controller"
)

func SetupRouter () *gin.Engine {
	r := gin.Default()
	
	r.GET("/pong", controller.Pong)

	return r
}