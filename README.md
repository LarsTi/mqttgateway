# MQTTGateway for Prometheus

A project that subscribes to MQTT queues and published prometheus metrics.
This Project is forked from [inuits/mqttgateway](https://github.com/inuits/mqttgateway).
It is designed to run in docker, especially in docker-compose.

## Prerequisites

docker and docker-compose should be installed correctly.


## Installation

````
git clone https://github.com/LarsTi/mqttgateway.git
cd mqttgateway
docker-compose up --build -d
````

## How does it work?

mqttgateway will connect to the MQTT broker at `MQTT_BROKER_ADDRESS` and
listen to the topics specified by `MQTT_TOPIC`.

The format for the topics is as follow:

`prefix/LABEL1/VALUE1/LABEL2/VALUE2/NAME`

A topic `prometheus/job/ESP8266/instance/livingroom/temperature_celcius` would
be converted to a metric
`temperature_celcius{job="ESP8266",instance="livingroom"}`.

If labelnames differ for a same metric, then we invalidate existing metrics and
only keep new ones. Then we issue a warning in the logs. You should avoid it.

Two other metrics are published, for each metric:

- `mqtt_NAME_last_pushed_timestamp`, the last time NAME metric has been pushed
(unix time, in seconds)
- `mqtt_NAME_push_total`, the number of times a metric has been pushed

## Available Environmentvariables:

`WEB_TELEMETRY_PORT` defaults to `9337` and is supposed to be the port where prometheus can scrape the metrics

`WEB_TELEMETRY_PATH` defaults to `/metrics` and is supposed to be the endpoint where prometheus can scrape

`MQTT_BROKER_ADDRESS` defaults to `tcp://localhost:1883` and is supposed to be the address to a MQTT-Broker

`MQTT_TOPIC` defaults to `prometheus/#` and is the topic which will be subscribed

`MQTT_PREFIX` defaults to `prometheus` and is the prefix which will be cut (see above)

`MQTT_CLIENT_ID` defaults to `mqtt2prometheus` and is the client ID which is used to identify to mqtt (not the username!)

## Not yet done

The Authentication towards the MQTT-Broker is not yet done and not planned. Due to my setup, the MQTT-Broker is in a seperated VLAN and therefore i do not use authentications on the broker.

## Example
This example is also in docker-compose.yaml

````
version: '3.3'

services:
        mqttgateway:
                build: .
                restart: unless-stopped

networks:
        default:

````

## A note about the prometheus config

If you use `job` and `instance` labels, please refer to the [pushgateway
exporter
documentation](https://github.com/prometheus/pushgateway#about-the-job-and-instance-labels).

TL;DR: you should set `honor_labels: true` in the scrape config.
