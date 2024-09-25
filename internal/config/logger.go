package config

import (
	"log/slog"
	"os"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

func SetUpLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case EnvLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case EnvDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case EnvProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return logger
}
