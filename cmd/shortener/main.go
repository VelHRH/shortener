package main

import (
	"os"
	"shortener/config"
	"shortener/lib/logger/sl"
	"shortener/storage/mysql"

	"golang.org/x/exp/slog"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger()
	log.Info("starting app", slog.String("env", cfg.Env))
	storage, err := mysql.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	log.Info("Stogage init", storage)
}

func setupLogger() *slog.Logger {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return log
}
