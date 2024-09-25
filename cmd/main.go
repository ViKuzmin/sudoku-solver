package main

import (
	"context"
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
		logger.Info("system call", osCall)
		server.Stop()
		cancel()
	}()

	server.Start()
}
