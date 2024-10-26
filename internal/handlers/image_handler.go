package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sudoku-solver/internal/image_processing/image_processor"
	"sudoku-solver/internal/solver/sudoku_solver"
	"time"
)

type ImageHandler struct {
	Logger    *slog.Logger
	Processor *image_processor.ImageProcessorV1
	Solver    *sudoku_solver.Solver
}

func NewImageHandler(logger *slog.Logger, processor *image_processor.ImageProcessorV1) *ImageHandler {
	solver := sudoku_solver.NewSolver(logger)
	return &ImageHandler{
		Logger:    logger,
		Processor: processor,
		Solver:    solver,
	}
}

func (processor *ImageHandler) GetAndroidShellScript(w http.ResponseWriter, r *http.Request) {
	logger := processor.Logger
	logger.Info("start process image")
	now := time.Now()
	battlefield := processor.Processor.GetBattlefield(r)

	if battlefield == "" {
		logger.Error("failed to decode image from request")
		http.Error(w, "failed to process file", http.StatusInternalServerError)
		return
	}

	solve, err := processor.Solver.GetScript(battlefield)

	if err != nil {
		logger.Error("failed to solve sudoku")
		http.Error(w, "failed to process file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, solve)
	logger.Info(fmt.Sprintf("finish process image. Time: %d ms", time.Now().Sub(now).Milliseconds()))
}

func (processor *ImageHandler) GetRawAnswerData(w http.ResponseWriter, r *http.Request) {
	logger := processor.Logger
	logger.Info("start process image")
	now := time.Now()

	battlefield := processor.Processor.GetBattlefield(r)

	if battlefield == "" {
		logger.Error("failed to decode image from request")
		http.Error(w, "failed to process file", http.StatusInternalServerError)
		return
	}

	solveSudoku := processor.Solver.SolveSudoku(battlefield)

	if len(solveSudoku) == 0 {
		logger.Error("failed to decode image from request")
		http.Error(w, "failed to process file", http.StatusInternalServerError)
		return
	}

	marshal, err := json.Marshal(solveSudoku)

	if err != nil {
		logger.Error("failed to marshal response")
		http.Error(w, "failed to process file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(marshal))
	logger.Info(fmt.Sprintf("finish process image. Time: %d ms", time.Now().Sub(now).Milliseconds()))
}
