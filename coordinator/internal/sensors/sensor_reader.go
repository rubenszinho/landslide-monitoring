package sensors

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func PublishSensorData(client mqtt.Client) {
	for {
		sensorType := "soil"
		sensorID := "1"
		measurementType := "moisture"

		value, err := ReadSensor(sensorType, sensorID, measurementType)
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

func ReadSensor(sensorType, sensorID, measurementType string) (float64, error) {
	// TODO(rubenszinho): Implement actual sensor reading logic
	return 0, fmt.Errorf("ReadSensor not implemented")
}
