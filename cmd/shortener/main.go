package main

import (
	"os"
	"shortener/config"
	mwLogger "shortener/http-server/middleware/logger"
	"shortener/lib/logger/sl"
	"shortener/storage/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger()

	storage, err := mysql.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	log.Info("Stogage init", storage)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
}

func setupLogger() *slog.Logger {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return log
}
