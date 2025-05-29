package config

import "os"

type Config struct {
	NatsURL      string
	SMTPServer   string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
}

func Load() *Config {
	return &Config{
		NatsURL:      os.Getenv("NATS_URL"),
		SMTPServer:   os.Getenv("SMTP_SERVER"),
		SMTPPort:     os.Getenv("SMTP_PORT"),
		SMTPUser:     os.Getenv("SMTP_USER"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
	}
}
