package main

import (
	"gomyloader/internal/benchmark"
	"gomyloader/internal/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	cfg.ShowConfig()

	benchmark.RunBenchmark(cfg)
}
