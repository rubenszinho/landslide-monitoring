package app

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rubenszinho/landslide-monitoring/internal/control"
)

func SubscribeControlTopics(client mqtt.Client) {
	wakeupTopic := "/control/wakeup/communication_unit/#"
	if token := client.Subscribe(wakeupTopic, 1, control.WakeupHandler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	restartTopic := "/control/restart/coordinator/#"
	if token := client.Subscribe(restartTopic, 1, control.RestartHandler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	shutdownTopic := "/control/shutdown/coordinator/#"
	if token := client.Subscribe(shutdownTopic, 1, control.ShutdownHandler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}
