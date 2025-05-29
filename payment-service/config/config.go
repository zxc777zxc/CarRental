package config

import (
	"os"
)

type Config struct {
	DBUrl   string
	NatsURL string
}

func Load() *Config {
	return &Config{
		DBUrl:   os.Getenv("DB_URL"),
		NatsURL: os.Getenv("NATS_URL"),
	}
}
