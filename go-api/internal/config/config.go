package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the .env file.
// If .env is not found, it will fallback to system environment variables.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, relying on system environment variables")
	}

	// Optional: check for required env vars here and log fatal if missing
	requiredVars := []string{"PORT"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Printf("Warning: environment variable %s is not set", v)
		}
	}
}
