package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
}

type Server struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"ssl_mode"`
}

func LoadConfig() (*Config, error) {
	cfg_path := os.Getenv("CONFIG_PATH")
	if cfg_path == "" {
		return nil,errors.New("CONFIG_PATH is not provided")
	}

	content, err := os.ReadFile(cfg_path)
	if err != nil {
		log.Fatal(err)
	}

	var	cfg *Config

	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return nil, errors.New("coulndt parse the config")
	}
	
	return cfg, nil
}