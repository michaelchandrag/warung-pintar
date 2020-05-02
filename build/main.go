package main

import (
	"os"
	"fmt"

	router "github.com/michaelchandrag/warung-pintar/module/router"
	websocket "github.com/michaelchandrag/warung-pintar/utils/websocket"
)

func main () {
	// initiate websocket
	websocket.WsHub = websocket.NewHub()
	go websocket.WsHub.Run()

	// initiate gin routing
	r := router.SetupRouter()

	// initiate running port
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}