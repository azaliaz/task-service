package storage

import (
	"log/slog"
	"time"
)

// Config хранит конфигурацию сервиса задач
type Config struct {
	TaskDuration time.Duration

	Logger *slog.Logger
}

func NewDefaultConfig() *Config {
	return &Config{
		TaskDuration: 4 * time.Minute,
		Logger:       slog.Default(),
	}
}
