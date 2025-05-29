package config

import "os"

type Config struct {
	PostgresDSN string
	RedisAddr   string
	GRPCPort    string
}

func Load() *Config {
	return &Config{
		PostgresDSN: getEnv("POSTGRES_DSN", "postgres://postgres:123@postgres:5432/car_rental?sslmode=disable"),
		RedisAddr:   getEnv("REDIS_ADDR", "localhost:6379"),
		GRPCPort:    getEnv("GRPC_PORT", ":50053"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
