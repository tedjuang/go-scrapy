package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config holds the application configuration
type Config struct {
	Server struct {
		Port int    `json:"port"`
		Host string `json:"host"`
	} `json:"server"`

	Data struct {
		Dir  string `json:"dir"`
		File string `json:"file"`
	} `json:"data"`

	Scraping struct {
		UserAgent string `json:"userAgent"`
		Timeout   int    `json:"timeout"`
		Retries   int    `json:"retries"`
	} `json:"scraping"`

	API struct {
		RateLimit  int `json:"rateLimit"`
		MaxResults int `json:"maxResults"`
	} `json:"api"`
}

// LoadConfig loads the configuration from the specified environment
func LoadConfig(env string) (*Config, error) {
	if env == "" {
		env = "dev" // Default to dev environment
	}

	configPath := filepath.Join("configs", env, "config.json")

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Parse config file
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return &config, nil
}
