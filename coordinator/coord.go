package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

var client mqtt.Client

func init() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading config.env file")
	}
}

func main() {
	broker := os.Getenv("MQTT_SERVER")
	clientID := "coordinator"

	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
	opts.SetCleanSession(false) // Guarantee persistence
	opts.SetDefaultPublishHandler(messageHandler)
	client = mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	go publishSensorData()
	go publishHealthMetrics()
	subscribeControlTopics()

	// Keep the main routine running
	select {}
}

func messageHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func publishSensorData() {
	for {
		sensorType := "soil"
		sensorID := "1"
		measurementType := "moisture"

		value, err := readSensor(sensorType, sensorID, measurementType)
		if err != nil {
			log.Printf("Error reading sensor data: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		topic := fmt.Sprintf("/data/coordinator/sensor/%s/%s/%s", sensorType, sensorID, measurementType)
		payload := fmt.Sprintf("%v", value)
		token := client.Publish(topic, 1, false, payload)
		token.Wait()
		fmt.Printf("Published sensor data to topic %s: %s\n", topic, payload)
		time.Sleep(10 * time.Second)
	}
}

func readSensor(sensorType, sensorID, measurementType string) (float64, error) {
	// TODO(rubenszinho): Implement actual sensor reading logic
	return 0, fmt.Errorf("readSensor not implemented")
}

func publishHealthMetrics() {
	for {
		cpuUsage, err := getCPUUsage()
		if err != nil {
			log.Printf("Error getting CPU usage: %v", err)
		} else {
			topicCPU := "/data/coordinator/health/cpu"
			payloadCPU := fmt.Sprintf("%f", cpuUsage)
			tokenCPU := client.Publish(topicCPU, 1, false, payloadCPU)
			tokenCPU.Wait()
			fmt.Printf("Published CPU usage to topic %s: %s\n", topicCPU, payloadCPU)
		}

		memoryUsage, err := getMemoryUsage()
		if err != nil {
			log.Printf("Error getting memory usage: %v", err)
		} else {
			topicMemory := "/data/coordinator/health/memory"
			payloadMemory := fmt.Sprintf("%f", memoryUsage)
			tokenMemory := client.Publish(topicMemory, 1, false, payloadMemory)
			tokenMemory.Wait()
			fmt.Printf("Published memory usage to topic %s: %s\n", topicMemory, payloadMemory)
		}

		time.Sleep(30 * time.Second)
	}
}

func getCPUUsage() (float64, error) {
	// TODO(rubenszinho): Implement actual CPU usage retrieval
	return 0, fmt.Errorf("getCPUUsage not implemented")
}

func getMemoryUsage() (float64, error) {
	// TODO(rubenszinho): Implement actual memory usage retrieval
	return 0, fmt.Errorf("getMemoryUsage not implemented")
}

func subscribeControlTopics() {
	wakeupTopic := "/control/wakeup/communication_unit/#"
	if token := client.Subscribe(wakeupTopic, 1, wakeupHandler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	restartTopic := "/control/restart/coordinator/#"
	if token := client.Subscribe(restartTopic, 1, restartHandler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	shutdownTopic := "/control/shutdown/coordinator/#"
	if token := client.Subscribe(shutdownTopic, 1, shutdownHandler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}

func wakeupHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Wake-up command received: %s from topic: %s\n", msg.Payload(), msg.Topic())
	// TODO(rubenszinho): Implement logic to handle wake-up commands for communication units
}

func restartHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Restart command received: %s from topic: %s\n", msg.Payload(), msg.Topic())
	topicLevels := splitTopic(msg.Topic())

	if len(topicLevels) < 5 {
		fmt.Println("Invalid restart command topic.")
		return
	}
	restartType := topicLevels[4]

	switch restartType {
	case "soft":
		fmt.Println("Performing soft restart...")
		softRestart()
	case "full":
		fmt.Println("Performing full restart...")
		fullRestart()
	default:
		fmt.Printf("Unknown restart type: %s\n", restartType)
	}
}

func shutdownHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Shutdown command received: %s from topic: %s\n", msg.Payload(), msg.Topic())
	topicLevels := splitTopic(msg.Topic())

	if len(topicLevels) < 5 {
		fmt.Println("Invalid shutdown command topic.")
		return
	}
	shutdownType := topicLevels[4]

	switch shutdownType {
	case "graceful":
		fmt.Println("Performing graceful shutdown...")
		gracefulShutdown()
	case "immediate":
		fmt.Println("Performing immediate shutdown...")
		immediateShutdown()
	default:
		fmt.Printf("Unknown shutdown type: %s\n", shutdownType)
	}
}

func splitTopic(topic string) []string {
	return strings.Split(topic, "/")
}

func softRestart() {
	fmt.Println("Soft restart not implemented yet.")
	// TODO(rubenszinho): Implement soft restart logic
}

func fullRestart() {
	fmt.Println("Attempting full system restart...")
	if runtime.GOOS == "linux" {
		cmd := exec.Command("sudo", "reboot")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Failed to reboot:", err)
		}
	} else {
		fmt.Println("Full restart is not supported on this OS.")
	}
}

func gracefulShutdown() {
	fmt.Println("Graceful shutdown initiated.")
	client.Disconnect(250)
	os.Exit(0)
}

func immediateShutdown() {
	fmt.Println("Immediate shutdown initiated.")
	client.Disconnect(250)
	os.Exit(1)
}
