package app

import (
	"fmt"
	"time"

	"github.com/rubenszinho/landslide-monitoring/internal/sensors"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func PublishSensorData(mqttClient mqtt.Client) {
	for {
		sensorType := SoilSensor
		sensorID := "1"
		measurementType := Salinity

		value, err := sensors.ReadSensor(sensorType, sensorID, measurementType)
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}

		topic := fmt.Sprintf("/data/coordinator/sensor/%s/%s/%s",
			sensorType.String(), sensorID, measurementType.String())

		payload := fmt.Sprintf("%v", value)
		token := mqttClient.Publish(topic, 1, false, payload)
		token.Wait()

		time.Sleep(10 * time.Second)
	}
}
