package internal

import (
	"context"
	"log/slog"
	"net/http"
	"sudoku-solver/api"
	"sudoku-solver/internal/config"
	"sudoku-solver/internal/handlers"
	"sudoku-solver/internal/image_processing/image_processor"
	"sync"
	"time"
)

type Server struct {
	logger       *slog.Logger
	ctx          context.Context
	cfg          *config.EnvironmentConfig
	Server       *http.Server
	ImageHandler *handlers.ImageHandler
	wg           *sync.WaitGroup
}

func NewServer(logger *slog.Logger, ctx context.Context, cfg *config.EnvironmentConfig) *Server {
	processor := image_processor.NewImageProcessorV1(logger)
	imgHandler := handlers.NewImageHandler(logger, processor)
	wg := new(sync.WaitGroup)

	return &Server{
		logger:       logger,
		ctx:          ctx,
		cfg:          cfg,
		ImageHandler: imgHandler,
		wg:           wg,
	}
}

func (server *Server) Start() {
	routes := api.CreateRoutes(server.ImageHandler)
	server.logger.LogAttrs(
		server.ctx,
		slog.LevelInfo,
		"starting server with",
		slog.Any("config", server.cfg),
	)

	server.Server = &http.Server{
		Addr:    ":" + server.cfg.ServerConfig.Port,
		Handler: routes,
	}

	server.wg.Add(1)
	err := server.Server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		server.logger.LogAttrs(
			server.ctx,
			slog.LevelError,
			"failed to listen and serve",
			slog.Any("err", err),
		)
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
		server.logger.LogAttrs(
			server.ctx,
			slog.LevelError,
			"server shutdown failed",
			slog.Any("err", err),
		)
	}

	server.logger.Info("server stopped")
	server.wg.Done()
}
