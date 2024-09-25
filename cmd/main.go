package main

import (
	"context"
	"ocr-test/internal"
	"ocr-test/internal/config"
	"os"
	"os/signal"
)

func main() {
	cfg := config.MustLoadEnvironmentConfig()
	logger := config.SetUpLogger(cfg.ServerConfig.Env)

	_, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	server := internal.NewServer(logger, cfg)

	go func() {
		osCall := <-c
		logger.Info("system call", osCall)
		server.Stop()
		cancel()
	}()

	server.Start()
}
