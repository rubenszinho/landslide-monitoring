#!/bin/bash

# Start Mosquitto in the background
/usr/sbin/mosquitto -c /mosquitto/config/mosquitto.conf &

# Wait a few seconds to ensure Mosquitto is up and running
sleep 5

# Start the Go application
/usr/local/bin/coordinator
