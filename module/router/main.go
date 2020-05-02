package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/michaelchandrag/warung-pintar/module/controller"
)

func SetupRouter () *gin.Engine {
	r := gin.Default()
	
	// load template views and assets
	r.LoadHTMLGlob("public/views/*")
	r.Static("/js", "public/assets/js")

	// load router
	r.GET("/", controller.ServePage)
	r.POST("/send", controller.Send)
	r.GET("/messages", controller.Find)
	
	// load websocket
	r.GET("/ws", controller.ServeWs)

	return r
}