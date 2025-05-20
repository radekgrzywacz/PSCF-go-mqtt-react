package sensor

import (
	"go-api/internal/mqtt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	mqtt mqtt.Client
}

type temperatureResponse struct {
	Temperature int    `json:"temperature"`
	Source      string `json:"source"`
}

func NewHandler(mqtt *mqtt.Client) *Handler {
	return &Handler{
		mqtt: *mqtt,
	}
}

func (h *Handler) GetTemperature(c *gin.Context) {
	c.JSON(http.StatusOK, temperatureResponse{
		Temperature: rand.Intn(100),
		Source:      "esp32",
	})
}

type humidityResponse struct {
	Humidity int    `json:"humidity"`
	Source   string `json:"source"`
}

func (h *Handler) GetHumidity(c *gin.Context) {
	c.JSON(http.StatusOK, humidityResponse{
		Humidity: rand.Intn(100),
		Source:   "esp32",
	})
}
