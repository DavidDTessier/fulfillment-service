package main

import (
	"os"

	service "github.com/DuffMan2020/fulfillment-service/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3001"
	}

	server := service.NewServer()
	server.Run(":" + port)
}
