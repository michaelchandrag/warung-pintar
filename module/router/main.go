package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/michaelchandrag/warung-pintar/module/controller"
)

func SetupRouter () *gin.Engine {
	r := gin.Default()
	
	r.GET("/", controller.HelloWorld)
	r.POST("/send", controller.Send)
	r.GET("/messages", controller.Find)

	return r
}