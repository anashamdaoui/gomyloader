# Go Registry Benchmark Tool

This benchmark tool is designed to load test the registry service by sending configurable HTTP requests to its endpoints. It supports different load profiles and allows you to monitor the performance using Prometheus.

## Features
- **Configurable Endpoints**: Define endpoints, HTTP methods, and payloads in `config.yaml`.
- **Load Profiles**: Supports fixed, ramp up, spike, and peak load profiles.
- **Virtual Users**: Specify the number of virtual users (simultaneous requests).
- **Test Duration**: Set test duration in minutes per endpoint.
- **Metrics Export**: Exports metrics to Prometheus for analysis.

## Monitoring

	•	The tool exposes metrics on http://localhost:2112/metrics.
	•	Use Prometheus to scrape these metrics for analysis.

## Load Profiles

### Fixed:
    •	Maintains a constant number of virtual users throughout the test.
	•	Parameters:
```yaml
    load_profile: "fixed"
    max_load: <maximum_number_of_virtual_users>
    duration_minutes: <test_duration_in_minutes>
```

### Ramp Up:
    •	Gradually increases the number of virtual users from an initial load to the maximum load.
	•	Parameters:
```yaml
    load_profile: "ramp_up"
    initial_load: <initial_number_of_virtual_users>
    max_load: <maximum_number_of_virtual_users>
    duration_minutes: <test_duration_in_minutes>
```

### Spike:
    •	Brief spike to maximum load, then back to base load.
	•	Parameters:
```yaml
    load_profile: "spike"
    base_load: <base_number_of_virtual_users>
    max_load: <maximum_number_of_virtual_users>
    duration_minutes: <test_duration_in_minutes>
```

### Peak:
    •	Increases to peak load, holds, then decreases back to base load.
	•	Parameters:
```yaml
    load_profile: "peak"
    base_load: <base_number_of_virtual_users>
    max_load: <maximum_number_of_virtual_users>
    duration_minutes: <test_duration_in_minutes>
```

## Getting Started

### Prerequisites
- Go 1.22 or later
- Docker (optional, for Prometheus setup)

### Installation
1. Clone the repository:
```bash
   git clone https://github.com/ahamdaoui/benchmark_tool.git
```
2. Navigate to the project folder
```bash 
cd benchmark_tool
```
3. Install dependencies
```bash
    go mod tidy
````
4. Edit config/config.yaml to specify endpoints, load profiles, and other settings:
```yaml
    endpoints:
        - path: "/register"
    method: "POST"
    payloads: 
      - '{"id": "worker1", "port": 8080}'
    load_profile: "ramp_up"
    initial_load: 1
    max_load: 10
    duration_minutes: 5
    registry_base_url: "http://localhost:8080"
```
5. Run
```bash
    go run ./cmd/main.go
```

