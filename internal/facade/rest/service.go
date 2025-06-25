package rest

import (
	"context"
	"fmt"
	"github.com/azaliaz/task-service/internal/application"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"time"
)

type Service struct {
	log    *slog.Logger
	config *Config
	fiber  *fiber.App
	app    application.TaskService
}

func NewAPI(
	logEntry *slog.Logger,
	config *Config,
	app application.TaskService,
) *Service {
	return &Service{
		log:    logEntry,
		config: config,
		app:    app,
	}
}

func (api *Service) Init() error {
	api.fiber = fiber.New(fiber.Config{
		ReadTimeout:           time.Duration(api.config.FiberReadTimeout) * time.Second,
		WriteTimeout:          time.Duration(api.config.FiberWriteTimeout) * time.Second,
		IdleTimeout:           time.Duration(api.config.FiberIdleTimeout) * time.Second,
		BodyLimit:             int(api.config.FiberBodyLimit),
		ReadBufferSize:        int(api.config.FiberReadBufferSize),
		StrictRouting:         api.config.FiberStrictRouting,
		CaseSensitive:         api.config.FiberCaseSensitive,
		DisableStartupMessage: api.config.FiberDisableStartupMessage,
		DisableKeepalive:      api.config.FiberDisableKeepalive,
	})

	api.fiber.Post("/tasks", api.CreateTask)
	api.fiber.Get("/tasks/:id", api.GetTaskStatus)
	api.fiber.Delete("/tasks/:id", api.DeleteTask)

	return nil
}

func (api *Service) Run(ctx context.Context) {
	addr := fmt.Sprintf(":%d", api.config.Port)
	api.log.Info("start rest server", "addr", addr)
	if err := api.fiber.Listen(addr); err != nil {
		api.log.Error("failed to start rest server", "addr", addr, "error", err)
	}
}

func (api *Service) Stop() {

}
