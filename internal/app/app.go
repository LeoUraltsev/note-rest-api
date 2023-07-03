package app

import (
	"NoteRestApi/config"
	"golang.org/x/exp/slog"
	stdLog "log"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Run() {

	cfg, err := config.Load()
	if err != nil {
		stdLog.Fatal(err)
	}

	log := setupLogger(cfg.Env)

	log.Info("starting note service", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// TODO: init storage: postgresql

	// TODO: init router: chi

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
