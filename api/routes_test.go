package api

import (
	"github.com/gorilla/mux"
	"log/slog"
	"os"
	"sudoku-solver/internal/handlers"
	"sudoku-solver/internal/image_processing/image_processor"
	"testing"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

func TestCreateRoutes(t *testing.T) {
	type args struct {
		imageHandler *handlers.ImageHandler
	}

	processorV1 := image_processor.NewImageProcessorV1(logger)
	handler := handlers.NewImageHandler(logger, processorV1)

	arg := args{
		imageHandler: handler,
	}

	tests := []struct {
		name string
		args args
		want *mux.Router
	}{
		{
			name: "test_1",
			args: arg,
			want: CreateRoutes(handler),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateRoutes(tt.args.imageHandler)
			if got == nil {
				t.Errorf("CreateRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
