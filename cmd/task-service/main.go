package main

import (
	"context"
	"flag"
	"github.com/azaliaz/task-service/internal/application"
	"github.com/azaliaz/task-service/internal/facade/rest"
	"github.com/azaliaz/task-service/pkg/config"
	"github.com/azaliaz/task-service/pkg/service"
	"log/slog"
	"os"
)

type Config struct {
	App  application.Config `envPrefix:"APP_" yaml:"app"`
	Rest rest.Config        `envPrefix:"REST_" yaml:"rest"`
}

func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	configFile := flag.String("config-file", "none", "path to config file")
	flag.Parse()

	cfg := Config{}
	if err := config.ReadConfig(*configFile, &cfg); err != nil {
		logger.Error("config parse error", "err_msg", err)
		os.Exit(1)
	}

	app := application.NewService(logger, &cfg.App)

	restConfig := cfg.Rest
	api := rest.NewAPI(logger, &restConfig, app)

	mgr := service.NewManager(logger)
	mgr.AddService(app, api)

	ctx := context.Background()
	if err := mgr.Run(ctx); err != nil {
		logger.Error("failed to start services", "err", err)
		os.Exit(1)
	}
}
