package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sudoku-solver/internal"
	"sudoku-solver/internal/config"
)

func main() {
	cfg := config.MustLoadEnvironmentConfig()
	logger := config.SetUpLogger(cfg.ServerConfig.Env)

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	server := internal.NewServer(logger, ctx, cfg)

	go func() {
		osCall := <-c
		logger.Info("system call", slog.Any("signal", osCall))
		server.Stop()
		cancel()
	}()

	server.Start()
}
