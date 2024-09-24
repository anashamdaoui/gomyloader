package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type EndpointConfig struct {
	Path            string            `yaml:"path"`
	Method          string            `yaml:"method"`
	Payloads        []string          `yaml:"payload"`
	Params          string            `yaml:"params,omitempty"`
	Headers         map[string]string `yaml:"headers,omitempty"`
	LoadProfile     string            `yaml:"load_profile"`
	BaseLoad        int               `yaml:"base_load,omitempty"`
	MaxLoad         int               `yaml:"max_load"`
	DurationMinutes int               `yaml:"duration_minutes"`
}

type Config struct {
	Endpoints       []EndpointConfig `yaml:"endpoints"`
	RegistryBaseURL string           `yaml:"registry_base_url"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (cfg *Config) ShowConfig() {
	log.Println("Loaded Configuration:")
	log.Printf("Registry Base URL: %s\n", cfg.RegistryBaseURL)
	for _, endpoint := range cfg.Endpoints {
		fmt.Printf("Endpoint Path: %s\n", endpoint.Path)
		fmt.Printf("Method: %s\n", endpoint.Method)
		fmt.Printf("Payloads: %v\n", endpoint.Payloads)
		fmt.Printf("Headers: %v\n", endpoint.Headers)
		fmt.Printf("Load Profile: %s\n", endpoint.LoadProfile)
		fmt.Printf("Base Load: %d\n", endpoint.BaseLoad)
		fmt.Printf("Max Load: %d\n", endpoint.MaxLoad)
		fmt.Printf("Duration (minutes): %d\n", endpoint.DurationMinutes)
		fmt.Println("---")
	}
}
