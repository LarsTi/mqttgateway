package main

import (
	"net/http"
	"os"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
//	"gopkg.in/alecthomas/kingpin.v2"
)
func getEnv(key, fallback string) string {
	log.Infoln(key)
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
var (
	listenAddress = ":9337"
	metricsPath = getEnv("WEB_TELEMETRY_PATH","/metrics")
	brokerAddress = getEnv("MQTT_BROKER_ADDRESS","tcp://localhost:1883")
	topic = getEnv("MQTT_TOPIC","prometheus/#")
	prefix = getEnv("MQTT_PREFIX","prometheus")
	mqttClientId = getEnv("MQTT_ID","mqtt2prometheus")
	progname  = "mqttgateway"
)

func main() {
//	log.AddFlags(kingpin.CommandLine)
//	kingpin.Parse()

	prometheus.MustRegister(newMQTTExporter())

	http.Handle(metricsPath, promhttp.Handler())
	log.Infoln("Listening on", listenAddress)
	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
