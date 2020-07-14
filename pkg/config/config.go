package config

import (
	"encoding/json"
	"os"
)

// Config main app configuration struct.
type Config struct {
	DB struct {
		Path string `json:"path"`
		Init string `json:"init"`
	}
}

// LoadConfig creates a struct from file.
func LoadConfig(filepath string) (Config, error) {
	var config Config
	// Open file.
	file, err := os.Open("config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()

	// Parse config.
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
