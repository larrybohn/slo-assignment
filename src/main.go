package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var searchDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "search_duration_seconds",
		Help:    "Duration of search requests in seconds",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"type"},
)

func init() {
	prometheus.MustRegister(searchDuration)
}

type PageData struct {
	Hostname string
}

type Product struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

var tmpl, _ = template.ParseFiles("index.html")

func main() {

	http.HandleFunc("/", handleIndex)

	http.HandleFunc("/search", handleSearch)

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
