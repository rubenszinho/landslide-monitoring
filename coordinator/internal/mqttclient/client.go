package mqttclient

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewClient(broker, clientID string) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
	opts.SetCleanSession(false) // Guarantee persistence

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	return client
}
