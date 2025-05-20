package mqtt

import (
	"fmt"
	"go-api/internal/config"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	Client mqtt.Client
}

func NewMqttConfig(c *config.Config) *Client {
	broker := fmt.Sprintf("%s", c.MQTT.Broker)
	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetUsername(c.MQTT.Username).
		SetPassword(c.MQTT.Password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT connection error: %s", token.Error().Error())
	}

	return &Client{
		Client: client,
	}
}

func (c *Client) Publish(message, topic string) {
	token := c.Client.Publish(topic, 0, false, message)
	token.Wait()
}

func (c *Client) Subscribe(topic string, callback mqtt.MessageHandler) {
	if token := c.Client.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT connection error: %s", token.Error().Error())
	}
}
