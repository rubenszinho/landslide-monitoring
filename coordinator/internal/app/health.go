package app

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rubenszinho/landslide-monitoring/internal/health"
)

func PublishHealthMetrics(client mqtt.Client) {
	for {
		cpuUsage, err := health.GetCPUUsage()
		if err != nil {
			log.Printf("Error getting CPU usage: %v", err)
		} else {
			topicCPU := "/data/coordinator/health/cpu"
			payloadCPU := fmt.Sprintf("%f", cpuUsage)
			tokenCPU := client.Publish(topicCPU, 1, false, payloadCPU)
			tokenCPU.Wait()
			fmt.Printf("Published CPU usage to topic %s: %s\n", topicCPU, payloadCPU)
		}

		memoryUsage, err := health.GetMemoryUsage()
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
