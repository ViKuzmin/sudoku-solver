package internal

import (
	"context"
	"log/slog"
	"net/http"
	"ocr-test/api"
	"ocr-test/internal/config"
	"ocr-test/internal/handlers"
	"ocr-test/internal/image_processing/image_processor"
	"sync"
	"time"
)

type Server struct {
	logger       *slog.Logger
	cfg          *config.EnvironmentConfig
	Server       *http.Server
	ImageHandler *handlers.ImageHandler
	wg           *sync.WaitGroup
}

func NewServer(logger *slog.Logger, cfg *config.EnvironmentConfig) *Server {
	processor := image_processor.NewImageProcessorV1(logger)
	imgHandler := handlers.NewImageHandler(logger, processor)
	wg := new(sync.WaitGroup)

	return &Server{
		logger:       logger,
		cfg:          cfg,
		ImageHandler: imgHandler,
		wg:           wg,
	}
}

func (server *Server) Start() {
	routes := api.CreateRoutes(server.ImageHandler)
	server.logger.Info("starting server with config:", server.cfg)

	server.Server = &http.Server{
		Addr:    ":" + server.cfg.ServerConfig.Port,
		Handler: routes,
	}

	server.wg.Add(1)
	err := server.Server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		server.logger.Error("failed to listen and serve", err)
		server.wg.Done()
		panic("failed to start server")
	}

	server.wg.Wait()
}

func (server Server) Stop() {
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := server.Server.Shutdown(ctxShutDown); err != nil {
		server.logger.Error("server shutdown failed", err)
	}

	server.logger.Info("server stopped")
	server.wg.Done()
}
