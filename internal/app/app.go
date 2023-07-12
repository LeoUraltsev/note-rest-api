package app

import (
	"NoteRestApi/config"
	v1 "NoteRestApi/internal/controller/http/v1"
	mwLogger "NoteRestApi/internal/controller/http/v1/middleware/logger"
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

	log.Info("init database")
	pgUrl := os.Getenv("PG_CONN_PATH")
	if pgUrl == "" {
		log.Error("PG_CONN_PATH is not set")
		os.Exit(1)
	}
	client, err := postgres.NewClient(context.Background(), pgUrl)

	log.Info("create repositories")
	rep := repo.NewRepositories(client, log)

	log.Info("create services")
	services := service.NewNoteServices(rep)

	log.Info("init Chi router")
	router := chi.NewRouter()

	router.Use(mwLogger.New(log))

	log.Info("init router")
	v1.NewRouter(router, services)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.HttpServer.Host, cfg.HttpServer.Port),
		Handler: router,
	}

	log.Info("server starting", slog.String("Host", cfg.HttpServer.Host), slog.String("Port", cfg.HttpServer.Port))

	if err := srv.ListenAndServe(); err != nil {
		log.Error("Server starting error:", err)
		panic(err)
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
