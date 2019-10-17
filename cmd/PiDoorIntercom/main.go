package main

import (
	"github.com/FloThinksPi/PiDoorIntercom/internal/app/PiDoorIntercom/bell"
	routes2 "github.com/FloThinksPi/PiDoorIntercom/internal/app/PiDoorIntercom/routes"
	"os"
)

const defaultPort string = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	bell.StartBell()

	r := routes2.SetupRoutes()

	r.Run(":" + port)
}
