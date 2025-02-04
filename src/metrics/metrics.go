package metrics

import "github.com/prometheus/client_golang/prometheus"

var RequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "request_duration_seconds",
		Help:    "Duration of requests at the specified endpoint in seconds",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"endpoint"},
)

func RegisterMetrics() {
	prometheus.MustRegister(RequestDuration)
}
