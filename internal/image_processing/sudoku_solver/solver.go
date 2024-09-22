package sudoku_solver

import (
	"fmt"
	"ocr-test/internal/image_processing/script_creator"
)

const N = 9

type Solver struct {
}

func NewSolver() *Solver {
	return &Solver{}
}

func (solver *Solver) GetScript(data [][]int) string {
	if SolveSudoku(data) {
		fmt.Println("Решенная головоломка Sudoku:")
		printGrid(data)
	} else {
		fmt.Println("Невозможно найти решение для данной головоломки Sudoku.")
	}

	creator := script_creator.NewScriptCreator()

	return creator.GetScript(data)
}

// Проверка, что значение val может быть помещено в ячейку grid[row][col]
func isSafe(grid [][]int, row, col, val int) bool {
	// Проверяем строку и столбец
	for i := 0; i < N; i++ {
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
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

// Решаем головоломку Sudoku с помощью рекурсивного метода
func SolveSudoku(grid [][]int) bool {
	row, col := findEmptyLocation(grid)
	if row == -1 && col == -1 {
		return true // Если все ячейки заполнены, значит решение найдено
	}

	for num := 1; num <= 9; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num

			if SolveSudoku(grid) {
				return true
			}

			grid[row][col] = 0 // Backtrack
		}
	}

	return false
}

// Функция для вывода сетки Sudoku
func printGrid(grid [][]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%2d", grid[i][j])
		}
		fmt.Println()
	}
}
