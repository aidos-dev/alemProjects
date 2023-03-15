package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Http struct {
		Port         string `json:"port"`
		WriteTimeout int    `json:"write_timeout"`
		ReadTimeout  int    `json:"read_timeout"`
	} `json:"http"`
	Database struct {
		Name    string `json:"name"`
		Options string `json:"options"`
	} `json:"database"`
}

func InitConfig(path string) (*Config, error) {
	var cfg Config
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("init configs: open: %w", err)
	}
	if err = json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("init configs: decode: %w", err)
	}
	return &cfg, nil
}
