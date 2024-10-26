package sudoku_solver

import (
	"log/slog"
	"os"
	"reflect"
	"testing"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

var correctData = "....35...2.......7.....6....1.....5.4.4.7............7...87.........4...2.3.5......"
var unsolvableData = "....35...2.......7.....6....1.....5.4.4.7............7...87.........4...2.3.5......"

//var correctData = [][]int{
//	{0, 0, 0, 0, 4, 6, 0, 0, 0},
//	{3, 0, 0, 0, 0, 0, 0, 8, 0},
//	{0, 0, 0, 0, 7, 0, 0, 0, 0},
//	{2, 0, 0, 0, 0, 0, 6, 0, 5},
//	{0, 5, 0, 8, 0, 0, 0, 0, 0},
//	{0, 0, 0, 0, 0, 0, 7, 0, 0},
//	{0, 9, 7, 0, 0, 0, 0, 0, 0},
//	{0, 0, 0, 5, 0, 0, 0, 3, 0},
//	{4, 0, 6, 0, 0, 0, 0, 0, 0},
//}
//
//var unsolvableData = [][]int{
//	{1, 1, 0, 0, 4, 6, 0, 0, 0},
//	{3, 0, 0, 0, 0, 0, 0, 8, 0},
//	{0, 0, 0, 0, 7, 0, 0, 0, 0},
//	{2, 0, 0, 0, 0, 0, 6, 0, 5},
//	{0, 5, 0, 8, 0, 0, 0, 0, 0},
//	{0, 0, 0, 0, 0, 0, 7, 0, 0},
//	{0, 9, 7, 0, 0, 0, 0, 0, 0},
//	{0, 0, 0, 5, 0, 0, 0, 3, 0},
//	{4, 0, 6, 0, 0, 0, 0, 0, 0},
//}

//func TestSolveSudoku(t *testing.T) {
//	solver := NewSolver(logger)
//	type args struct {
//		grid [][]int
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		{
//			name: "test_1",
//			args: args{grid: correctData},
//			want: true,
//		},
//		{
//			name: "test_2",
//			args: args{grid: unsolvableData},
//			want: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			//if got := solver.SolveSudoku(tt.args.grid); got != tt.want {
//			if got := solver.SolveSudoku("tt.args.grid"); got != "" {
//				t.Errorf("SolveSudoku() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestSolver_GetScript(t *testing.T) {
//	type fields struct {
//		logger        *slog.Logger
//		scriptCreator *script_creator.ScriptCreator
//	}
//
//	field := fields{
//		logger:        logger,
//		scriptCreator: script_creator.NewScriptCreator(),
//	}
//
//	type args struct {
//		//data [][]int
//		data string
//	}
//	expectedScript := "input tap 90 505; input tap 889 2011; input tap 202 505; input tap 773 2011; input tap 314 505; input tap 77 2011; input tap 426 505; input tap 1005 2011; input tap 538 505; input tap 425 2011; input tap 650 505; input tap 657 2011; input tap 762 505; input tap 193 2011; input tap 874 505; input tap 541 2011; input tap 986 505; input tap 309 2011; input tap 130 619; input tap 309 2011; input tap 242 619; input tap 657 2011; input tap 354 619; input tap 425 2011; input tap 466 619; input tap 77 2011; input tap 578 619; input tap 193 2011; input tap 690 619; input tap 541 2011; input tap 802 619; input tap 1005 2011; input tap 914 619; input tap 889 2011; input tap 1026 619; input tap 773 2011; input tap 130 733; input tap 1005 2011; input tap 242 733; input tap 193 2011; input tap 354 733; input tap 541 2011; input tap 466 733; input tap 309 2011; input tap 578 733; input tap 773 2011; input tap 690 733; input tap 889 2011; input tap 802 733; input tap 77 2011; input tap 914 733; input tap 657 2011; input tap 1026 733; input tap 425 2011; input tap 130 847; input tap 193 2011; input tap 242 847; input tap 425 2011; input tap 354 847; input tap 889 2011; input tap 466 847; input tap 773 2011; input tap 578 847; input tap 77 2011; input tap 690 847; input tap 309 2011; input tap 802 847; input tap 657 2011; input tap 914 847; input tap 1005 2011; input tap 1026 847; input tap 541 2011; input tap 130 961; input tap 773 2011; input tap 242 961; input tap 541 2011; input tap 354 961; input tap 1005 2011; input tap 466 961; input tap 889 2011; input tap 578 961; input tap 657 2011; input tap 690 961; input tap 193 2011; input tap 802 961; input tap 309 2011; input tap 914 961; input tap 425 2011; input tap 1026 961; input tap 77 2011; input tap 130 1075; input tap 657 2011; input tap 242 1075; input tap 77 2011; input tap 354 1075; input tap 309 2011; input tap 466 1075; input tap 425 2011; input tap 578 1075; input tap 541 2011; input tap 690 1075; input tap 1005 2011; input tap 802 1075; input tap 773 2011; input tap 914 1075; input tap 193 2011; input tap 1026 1075; input tap 889 2011; input tap 130 1189; input tap 541 2011; input tap 242 1189; input tap 1005 2011; input tap 354 1189; input tap 773 2011; input tap 466 1189; input tap 657 2011; input tap 578 1189; input tap 309 2011; input tap 690 1189; input tap 425 2011; input tap 802 1189; input tap 889 2011; input tap 914 1189; input tap 77 2011; input tap 1026 1189; input tap 193 2011; input tap 130 1303; input tap 77 2011; input tap 242 1303; input tap 889 2011; input tap 354 1303; input tap 193 2011; input tap 466 1303; input tap 541 2011; input tap 578 1303; input tap 1005 2011; input tap 690 1303; input tap 773 2011; input tap 802 1303; input tap 425 2011; input tap 914 1303; input tap 309 2011; input tap 1026 1303; input tap 657 2011; input tap 130 1417; input tap 425 2011; input tap 242 1417; input tap 309 2011; input tap 354 1417; input tap 657 2011; input tap 466 1417; input tap 193 2011; input tap 578 1417; input tap 889 2011; input tap 690 1417; input tap 77 2011; input tap 802 1417; input tap 541 2011; input tap 914 1417; input tap 773 2011; input tap 1026 1417; input tap 1005 2011; "
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   string
//	}{
//		{
//			name:   "test_get_script",
//			fields: field,
//			args: args{
//				data: correctData,
//			},
//			want: expectedScript,
//		},
//		{
//			name:   "test_unsolved_data",
//			fields: field,
//			args: args{
//				data: unsolvableData,
//			},
//			want: "",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			solver := &Solver{
//				logger:        tt.fields.logger,
//				scriptCreator: tt.fields.scriptCreator,
//			}
//			if got, _ := solver.GetScript(tt.args.data); got != tt.want {
//				t.Errorf("GetScript() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestNewSolver(t *testing.T) {
	type args struct {
		logger *slog.Logger
	}
	tests := []struct {
		name string
		args args
		want *Solver
	}{
		{
			name: "test_new",
			args: args{
				logger: logger,
			},
			want: NewSolver(logger),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSolver(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSolver() = %v, want %v", got, tt.want)
			}
		})
	}
}
