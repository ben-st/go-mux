// this project is heavily inspired by the prometheus golang client example project
// https://github.com/prometheus/client_golang/blob/master/examples/simple/main.go

package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func main() {
	var log = logrus.New()

	// set the log output
	log.Out = os.Stdout

	// set the log level
	log.Level = logrus.DebugLevel

	// Routes consist of a path and a handler function.
	r := mux.NewRouter()

	// sample log.Info
	log.Info("http server is ready")

	// sample log.Debug
	log.Debug("i am only visible in debug mode\n")

	// exposes / endpoint with the YourHandler handler
	r.HandleFunc("/", YourHandler)

	// exposes /metrics endpoint with standard golang metrics used by prometheus
	r.Handle("/metrics", promhttp.Handler())

	// start a goroutine which start the polling for the metrics endpoint
	SimpleGauge()

	// wrap a logger around the mux server
	logWrapper := handlers.LoggingHandler(os.Stdout, r)

	metricsWrapper := prometheus.InstrumentHandler("metrics", logWrapper)

	// Bind to a port and pass our loggedRouter in
	log.Fatal(http.ListenAndServe(":8080", metricsWrapper))

}
