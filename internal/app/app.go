package app

import (
	"NoteRestApi/config"
	v1 "NoteRestApi/internal/controller/http/v1"
	"NoteRestApi/internal/repo"
	"NoteRestApi/internal/service"
	"NoteRestApi/pkg/postgres"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"
	stdLog "log"
	"net/http"
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

	client, err := postgres.NewClient(context.Background(), "postgresql://postgres:7892carat@localhost:5432/noterestapi")

	if err := client.Pool.Ping(context.Background()); err != nil {
		log.Error(fmt.Sprintf("Ping database error: %v", err))
	}

	rep := repo.NewRepositories(client)

	services := service.NewNoteServices(rep)

	handler := chi.NewRouter()

	v1.NewRouter(handler, services)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.HttpServer.Host, cfg.HttpServer.Port), handler)
	if err != nil {
		return
	}
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
