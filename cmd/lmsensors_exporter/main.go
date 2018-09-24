// Command lmsensors_exporter provides a Prometheus exporter for lmsensors
// sensor metrics.
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mdlayher/lmsensors"
	"github.com/mdlayher/lmsensors_exporter"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app				= kingpin.New("lmsensors_exporter", "Prometheus metrics exporter for lmsensors")
	telemetryAddr = app.Flag("web.listen-address", "Address to listen on for web interface and telemetry").Default(":9165").String()
	metricsPath   = app.Flag("web.telemetry-path","Path under which to expose metrics.").Default("/metrics").String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	prometheus.MustRegister(lmsensorsexporter.New(lmsensors.New()))

	http.Handle(*metricsPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, *metricsPath, http.StatusMovedPermanently)
	})

	log.Printf("starting lmsensors exporter on %q", *telemetryAddr)

	if err := http.ListenAndServe(*telemetryAddr, nil); err != nil {
		log.Fatalf("cannot start lmsensors exporter: %s", err)
	}
}
