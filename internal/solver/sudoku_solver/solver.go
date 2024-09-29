package sudoku_solver

import (
	"fmt"
	"log/slog"
	"sudoku-solver/internal/image_processing/script_creator"
)

const (
	n        = 9
	space    = " "
	nextLine = "\n"
)

type Solver struct {
	logger        *slog.Logger
	scriptCreator *script_creator.ScriptCreator
}

func NewSolver(logger *slog.Logger) *Solver {
	return &Solver{
		logger:        logger,
		scriptCreator: script_creator.NewScriptCreator(),
	}
}

func (solver *Solver) GetScript(data [][]int) (string, error) {
	if solver.SolveSudoku(data) {
		solver.logger.Info("successfully solved")
	} else {
		solver.logger.Info("failed to solve")
		return "", fmt.Errorf("failed to solve")
	}

	return solver.scriptCreator.GetScript(data), nil
}

// Проверка, что значение val может быть помещено в ячейку grid[row][col]
func isSafe(grid [][]int, row, col, val int) bool {
	// Проверяем строку и столбец
	for i := 0; i < n; i++ {
		if grid[row][i] == val || grid[i][col] == val {
			return false
		}
	}

	// Проверяем квадрат 3x3, в который входит ячейка grid[row][col]
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == val {
				return false
			}
		}
	}

	return true
}

// Находим пустую ячейку в сетке и возвращаем ее координаты
func findEmptyLocation(grid [][]int) (int, int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

// Решаем головоломку Sudoku с помощью рекурсивного метода
func (solver *Solver) SolveSudoku(grid [][]int) bool {
	row, col := findEmptyLocation(grid)
	if row == -1 && col == -1 {
		return true // Если все ячейки заполнены, значит решение найдено
	}

	for num := 1; num <= 9; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num

			if solver.SolveSudoku(grid) {
				return true
			}

			grid[row][col] = 0 // Backtrack
		}
	}

	return false
}
