package main

import (
	"go-api/internal/config"
	"go-api/internal/mqtt"
	"go-api/internal/routes"
	"log"
)

func main() {

	config := config.LoadConfig()
	mqtt := mqtt.NewMqttConfig(config)
	r := routes.SetupRoutes(mqtt)
	port := config.Port
	if port == "" {
		log.Fatal("No env PORT")
	}

	log.Printf("Server is running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server %v", err)
	}
}
