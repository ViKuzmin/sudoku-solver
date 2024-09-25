package handlers

import (
	"fmt"
	"image"
	"log/slog"
	"net/http"
	"sudoku-solver/internal/image_processing/image_processor"
	"sudoku-solver/internal/image_processing/sudoku_solver"
	"time"
)

const (
	fileKey = "file"
)

type ImageHandler struct {
	logger    *slog.Logger
	Processor *image_processor.ImageProcessorV1
	Solver    *sudoku_solver.Solver
}

func NewImageHandler(logger *slog.Logger, processor *image_processor.ImageProcessorV1) *ImageHandler {
	solver := sudoku_solver.NewSolver(logger)

	return &ImageHandler{
		logger:    logger,
		Processor: processor,
		Solver:    solver,
	}
}

func (processor *ImageHandler) ProcessImage(w http.ResponseWriter, r *http.Request) {
	logger := processor.logger
	logger.Info("start process image")
	now := time.Now()
	err := r.ParseMultipartForm(32 << 15)

	if err != nil {
		logger.Error("failed to parse data from request")
	}

	file, _, err := r.FormFile(fileKey)

	img, _, err := image.Decode(file)
	if err != nil {
		logger.Error("failed to decode image from request")
		http.Error(w, "failed to process file", http.StatusInternalServerError)
		return
	}

	battlefield := processor.Processor.GetBattlefield(img)
	solve := processor.Solver.GetScript(battlefield)

	fmt.Fprintf(w, solve)
	logger.Info(fmt.Sprintf("finish process image. Time: %s ms", time.Now().Sub(now).Milliseconds()))
}
