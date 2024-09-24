package benchmark

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_count",
			Help: "Total number of requests made to each endpoint",
		},
		[]string{"endpoint"},
	)

	requestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_latency_seconds",
			Help: "Latency of requests to each endpoint",
		},
		[]string{"endpoint"},
	)
)

func InitMetrics() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestLatency)
}

func RecordRequest(endpoint string, latency float64) {
	requestCount.WithLabelValues(endpoint).Inc()
	requestLatency.WithLabelValues(endpoint).Observe(latency)
}
