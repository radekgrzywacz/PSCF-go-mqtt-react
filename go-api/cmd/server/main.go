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
	
	log.Printf("Server is running on port %s", os.Getenv("PORT"))
	
	
}