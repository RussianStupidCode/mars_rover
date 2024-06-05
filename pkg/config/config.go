package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type Rover struct {
	ConnectionURL string  `env:"ROVER_CONNECTION_URL"`
	ErrPercent    float64 `env:"ROVER_ERR_PERCENT" default:"0.1"`
}

type Logger struct {
	Level string `env:"LOG_LEVEL" default:"DEBUG"`
}

type HealthChecker struct {
	Period time.Duration `env:"HEALTH_CHECKER_PERIOD" default:"5s"`
}

type Config struct {
	Rover         Rover
	Logger        Logger
	HealthChecker HealthChecker
}

func FromEnv() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("load config from env: %w", err)
	}
	return &cfg, nil
}
