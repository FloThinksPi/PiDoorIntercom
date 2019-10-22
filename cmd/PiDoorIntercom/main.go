package main

import (
	"github.com/FloThinksPi/PiDoorIntercom/internal/app/PiDoorIntercom/bell"
	routes2 "github.com/FloThinksPi/PiDoorIntercom/internal/app/PiDoorIntercom/routes"
	"os"
)

const defaultPort string = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	bell.InitBellWatcher()

	r := routes2.SetupRoutes()

	_ = r.Run(":" + port)
}
