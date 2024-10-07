package control

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rubenszinho/landslide-monitoring/internal/util"
)

func WakeupHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Wake-up command received: %s from topic: %s\n", msg.Payload(), msg.Topic())
	// TODO(rubenszinho): Implement logic to handle wake-up commands for communication units
}

func RestartHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Restart command received: %s from topic: %s\n", msg.Payload(), msg.Topic())
	topicLevels := util.SplitTopic(msg.Topic())

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

func ShutdownHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Shutdown command received: %s from topic: %s\n", msg.Payload(), msg.Topic())
	topicLevels := util.SplitTopic(msg.Topic())

	if len(topicLevels) < 5 {
		fmt.Println("Invalid shutdown command topic.")
		return
	}
	shutdownType := topicLevels[4]

	switch shutdownType {
	case "graceful":
		fmt.Println("Performing graceful shutdown...")
		gracefulShutdown(client)
	case "immediate":
		fmt.Println("Performing immediate shutdown...")
		immediateShutdown(client)
	default:
		fmt.Printf("Unknown shutdown type: %s\n", shutdownType)
	}
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

func gracefulShutdown(client mqtt.Client) {
	fmt.Println("Graceful shutdown initiated.")
	client.Disconnect(250)
	os.Exit(0)
}

func immediateShutdown(client mqtt.Client) {
	fmt.Println("Immediate shutdown initiated.")
	client.Disconnect(250)
	os.Exit(1)
}
