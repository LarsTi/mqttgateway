package main

import (
	"net/http"
	"os"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)
func getEnv(key, fallback string) string {
	log.Infoln(key)
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
var (
	listenAddress = getEnv("WEB_TELEMETRY_PORT","9337")
	metricsPath = getEnv("WEB_TELEMETRY_PATH","/metrics")
	brokerAddress = getEnv("MQTT_BROKER_ADDRESS","tcp://localhost:1883")
	topic = getEnv("MQTT_TOPIC","prometheus/#")
	prefix = getEnv("MQTT_PREFIX","prometheus")
	brokerClientId = getEnv("MQTT_CLIENT_ID","mqtt2prometheus")
	progname  = "mqttgateway"
)

func main() {
	prometheus.MustRegister(newMQTTExporter())

	http.Handle(metricsPath, promhttp.Handler())
	log.Infoln("Listening on port", listenAddress)
	err := http.ListenAndServe(":" + listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
