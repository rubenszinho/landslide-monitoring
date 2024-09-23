package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("config.env")
	println(os.Getenv("MQTT_SERVER"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var broker = os.Getenv("MQTT_SERVER")
var clientId = "coordinator"

func main() {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientId)
	opts.SetCleanSession(false) // Guarantee persistence
	opts.SetDefaultPublishHandler(messageHandler)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	go eventHandler(client)
	go handleData(client)
	go handleWakeUps(client)
	go dataFromUnits(client)

	// Keep the main routine running
	select {}
}

func messageHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func eventHandler(client mqtt.Client) {
	topic := "/Control/Event/#"
	if token := client.Subscribe(topic, 1, eventCallback); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}

func eventCallback(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Event received: %s from topic: %s\n", msg.Payload(), msg.Topic())

	instructionTopic := "/Data/Coordinator/"
	client.Publish(instructionTopic, 1, false, fmt.Sprintf("Instructions for event: %s", msg.Payload()))
}

func handleData(client mqtt.Client) {
	topic := "/Data/Coordinator/#"
	if token := client.Subscribe(topic, 1, dataCallback); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}

func dataCallback(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Data received: %s from topic: %s\n", msg.Payload(), msg.Topic())
	unitAwake := false
	if !unitAwake {
		wakeUpTopic := "/Control/WakeUp/"
		client.Publish(wakeUpTopic, 1, false, "Wake-up instructions")
		time.Sleep(2 * time.Second) // TODO: Implement actual wakeup signal
	}
	dataTopic := fmt.Sprintf("/Data/To/Unit/%s", msg.Payload())
	client.Publish(dataTopic, 1, false, "Forwarding data")
}

func handleWakeUps(client mqtt.Client) {
	topic := "/Control/WakeUp/#"
	if token := client.Subscribe(topic, 1, wakeUpCallback); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}

func wakeUpCallback(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Wake-up received: %s from topic: %s\n", msg.Payload(), msg.Topic())
	time.Sleep(2 * time.Second) // TODO: Implement actual wakeup signal
	fmt.Println("Unit awake")
}

func dataFromUnits(client mqtt.Client) {
	topic := "/Data/From/#"
	if token := client.Subscribe(topic, 1, dataFromUnitsCallback); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}

func dataFromUnitsCallback(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Data from unit received: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
