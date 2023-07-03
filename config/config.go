package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env        string     `yaml:"env" env-default:"local"`
	HttpServer HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-default:"8080"`
}

func Load() (*Config, error) {
	var cfg Config
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return &cfg, fmt.Errorf("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &cfg, fmt.Errorf("config file does not exist: %s", configPath)
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return &cfg, fmt.Errorf("cannot read config: %v", err)
	}

	return &cfg, nil

}
