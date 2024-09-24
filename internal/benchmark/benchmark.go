package benchmark

import (
	"gomyloader/internal/client"
	"gomyloader/internal/config"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/exp/rand"
)

func RunBenchmark(cfg *config.Config) {
	clientInstance := client.NewClient(cfg.RegistryBaseURL)
	InitMetrics()

	// Start Prometheus HTTP server
	go serveMetrics(":2112")

	var wg sync.WaitGroup // Add WaitGroup to track goroutines

	for _, endpoint := range cfg.Endpoints {
		log.Printf("Starting benchmark for endpoint %s with load profile: %s", endpoint.Path, endpoint.LoadProfile)

		loadProfile := GenerateLoadProfile(LoadProfile{
			Type:     endpoint.LoadProfile,
			BaseLoad: endpoint.BaseLoad,
			MaxLoad:  endpoint.MaxLoad,
			Duration: time.Duration(endpoint.DurationMinutes) * time.Minute,
		})

		end := time.Now().Add(time.Duration(endpoint.DurationMinutes) * time.Minute)
		for time.Now().Before(end) {
			for _, users := range loadProfile {
				log.Printf("Simulating %d virtual users for endpoint %s", users, endpoint.Path)
				for i := 0; i < users; i++ {
					wg.Add(1)                                  // Increment WaitGroup counter
					go simulate(&wg, clientInstance, endpoint) // Call simulate function
				}
				time.Sleep(1 * time.Second) // Control the load interval
			}
		}
		log.Printf("Completed benchmark for endpoint %s", endpoint.Path)
	}

	wg.Wait() // Wait for all goroutines to complete
	log.Println("Benchmark process completed for all endpoints.")
}

func simulate(wg *sync.WaitGroup, clientInstance *client.Client, ep config.EndpointConfig) {
	defer wg.Done() // Decrement counter when goroutine completes

	payload := ""
	if len(ep.Payloads) > 0 {
		payload = ep.Payloads[rand.Intn(len(ep.Payloads))]
	}

	resp, err := clientInstance.DoRequest(ep.Method, ep.Path, payload, ep.Params, ep.Headers)
	if err != nil {
		log.Printf("Request failed for endpoint %s: %v", ep.Path, err)
		return
	}
	defer resp.Body.Close()

	// Log request and response
	log.Printf("Request to %s with payload: %s", ep.Path, payload)
	log.Printf("Response Status: %s", resp.Status)

	// Read and log response body
	body, _ := io.ReadAll(resp.Body)
	log.Printf("Response Body: %s", string(body))

	RecordRequest(ep.Path, time.Since(time.Now()).Seconds())
	log.Printf("Request to %s completed", ep.Path)
}

var metricsOnce sync.Once

// ServeMetrics starts an HTTP server that exposes the Prometheus metrics endpoint
func serveMetrics(addr string) {
	metricsOnce.Do(func() {
		http.Handle("/metrics", promhttp.Handler()) // Expose the /metrics endpoint for Prometheus
		log.Printf("Starting registry metrics server on %s\n", addr)

		// Start the HTTP server to expose metrics
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("Error starting metrics server: %v", err)
		}
	})
}
