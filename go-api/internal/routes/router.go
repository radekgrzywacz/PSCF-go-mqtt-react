package routes

import (
	"go-api/internal/handler/sensor"
	"go-api/internal/mqtt"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(mqtt *mqtt.Client) *gin.Engine {
	r := gin.Default()
	sensorHandler := sensor.NewHandler(mqtt)

	api := r.Group("/api")
	{
		espData := api.Group("/data")
		{
			espData.GET("/temp", sensorHandler.GetTemperature)
			espData.GET("/humidity", sensorHandler.GetHumidity)
		}
	}

	return r
}
