#include <WiFi.h>
#include <PubSubClient.h>
#include "config.h"

const int mqtt_port = 1883;

const char* mqtt_topic_subscribe = "/Data/To/Unit/1"; 
const char* mqtt_topic_publish = "/Data/From/Unit/1";

WiFiClient espClient;
PubSubClient client(espClient);

void setup_wifi() {
  delay(10);
  Serial.println("Connecting to WiFi...");
  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }

  Serial.println("WiFi connected");
}

void connectToMQTT() {
  while (!client.connected()) {
    Serial.println("Connecting to MQTT...");
    if (client.connect("ESP32Client", NULL, NULL, mqtt_topic_publish, 1, false, "Offline")) {
      Serial.println("Connected to MQTT broker");
      client.subscribe(mqtt_topic_subscribe, 1);
    } else {
      Serial.print("Failed, rc=");
      Serial.print(client.state());
      Serial.println(" try again in 5 seconds");
      delay(5000);
    }
  }
}

void RelayData(const String& message) {
  Serial.println("Relaying data to outside world...");
  // TODO(samuel): Implement the whole logic behind data transmission to external world
}

void ToCoord(const String& data) {
  Serial.println("Sending data to coordinator...");
  client.publish(mqtt_topic_publish, data.c_str(), true);
}

void callback(char* topic, byte* payload, unsigned int length) {
  Serial.print("Message arrived on topic: ");
  Serial.print(topic);
  Serial.print(". Message: ");
  
  String message;
  for (int i = 0; i < length; i++) {
    message += (char)payload[i];
  }
  Serial.println(message);

  RelayData(message);

  ToCoord("Processed data: " + message);
}

void setup() {
  Serial.begin(115200);

  setup_wifi();

  client.setServer(mqtt_server, mqtt_port);
  client.setCallback(callback);

  connectToMQTT();

  client.subscribe(mqtt_topic_subscribe, 1);
}

void loop() {
  if (!client.connected()) {
    connectToMQTT();
  }
  client.loop();

  delay(1000);
}
