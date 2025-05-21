package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type MQTTConfig struct {
	Broker   string
	Port     string
	Username string
	Password string
}

type Config struct {
	Port string
	MQTT MQTTConfig
	Env  string
}

func LoadConfig() *Config {
	LoadEnv()

	cfg := &Config{
		Port: getEnv("PORT", "8080"),
		MQTT: MQTTConfig{
			Broker:   getEnv("MQTT_BROKER", "server.radekgrzywacz.pl"),
			Port:     getEnv("MQTT_PORT", "1883"),
			Username: getEnv("MQTT_USERNAME", "user"),
			Password: getEnv("MQTT_PASSWORD", "user"),
		},
		Env: getEnv("APP_ENV", "local"),
	}

	if cfg.MQTT.Broker == "" {
		log.Fatal("MQTT Broker must be set")
	}

	return cfg
}

// LoadEnv loads environment variables from the .env file.
// If .env is not found, it will fallback to system environment variables.
func LoadEnv() {
	err := godotenv.Load(".env.local")
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

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
