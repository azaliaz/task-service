package application

import (
	"log/slog"
	"time"
)

type Config struct {
	TaskDuration time.Duration `env:"TASK_DURATION" envDefault:"240s"`

	Secret string

	Logger *slog.Logger
}
