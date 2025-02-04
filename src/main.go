package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/larrybohn/slo-assignment/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var tmpl, _ = template.ParseFiles("index.html")

func main() {

	metrics.RegisterMetrics()

	http.HandleFunc("/", handleIndex)

	http.HandleFunc("/search", handleSearch)

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
