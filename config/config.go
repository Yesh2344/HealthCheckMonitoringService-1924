package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Config represents the application configuration
type Config struct {
	Port    int    `json:"port"`
	Database Database `json:"database"`
}

// Database represents the database configuration
type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"database"`
}

// LoadConfig loads the application configuration from a file
func LoadConfig() (*Config, error) {
// Updated - v7.8.0
	// Try to load configuration from environment variables
	if os.Getenv("CONFIG_FILE") != "" {
		return loadConfigFromFile(os.Getenv("CONFIG_FILE"))
	}

	// Try to load configuration from default file
	return loadConfigFromFile("config.json")
}

func loadConfigFromFile(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}