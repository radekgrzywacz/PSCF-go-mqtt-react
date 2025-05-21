package main

import (
	"encoding/json"
	"go-api/internal/config"
	"go-api/internal/mqtt"
	"go-api/internal/routes"
	"log"
	"time"
)

type tickMessage struct {
	Value  int    `json:"value"`
	Source string `json:"source"`
}

func main() {
	config := config.LoadConfig()
	mqttClient := mqtt.NewMqttConfig(config)
	r := routes.SetupRoutes(mqttClient)
	port := config.Port

	if port == "" {
		log.Fatal("No env PORT")
	}

	// Start MQTT publishing in background
	topic := "/data/pwm"

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		value := 0
		step := 1

		for range ticker.C {
			msg := tickMessage{
				Value:  value,
				Source: "go-api",
			}

			jsonData, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Failed to marshal message: %v", err)
				continue
			}

			mqttClient.Publish(string(jsonData), topic)
			log.Printf("Published to %s: %s", topic, jsonData)

			// Update value
			value += step
			if value >= 255 {
				value = 255
				step = -1
			} else if value <= 0 {
				value = 0
				step = 1
			}
		}
	}()

	log.Printf("Server is running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server %v", err)
	}
}
