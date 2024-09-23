#!/bin/bash

/usr/sbin/mosquitto -c /mosquitto/config/mosquitto.conf &

sleep 5

/usr/local/bin/coordinator
