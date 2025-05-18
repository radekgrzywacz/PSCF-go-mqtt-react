package main

import (
	"go-api/internal/config"
	"go-api/internal/routes"
	"log"
	"os"
)

func main() {
	config.LoadEnv()

	r := routes.SetupRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No env PORT")
	}

	log.Printf("Server is running on port %s", os.Getenv("PORT"))
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to run server %v", err)
	}

}
