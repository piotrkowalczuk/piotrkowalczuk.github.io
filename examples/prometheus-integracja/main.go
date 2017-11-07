package main

import (
	"io"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	app := http.NewServeMux()
	app.Handle("/", &handler{})
	http.ListenAndServe("0.0.0.0:8080", app)

	dbg := http.NewServeMux()
	dbg.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{},
	))
	http.ListenAndServe("0.0.0.0:8081", dbg)
}

type handler struct {
}

func (h *handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "To działa!")
}
