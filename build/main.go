package main

import (
	"os"
	"fmt"

	router "github.com/michaelchandrag/warung-pintar/module/router"
)

func main () {
	r := router.SetupRouter()

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}