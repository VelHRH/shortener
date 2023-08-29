package main

import (
	"os"
	"shortener/config"

	"golang.org/x/exp/slog"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger()
	log.Info("starting app", slog.String("env", cfg.Env))
}

func setupLogger() *slog.Logger {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return log
}
