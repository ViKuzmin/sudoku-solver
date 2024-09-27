package image_processor

import (
	"errors"
	"log/slog"
	"os"
	"reflect"
	"testing"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
var sample = "sample.jpg"

func BenchmarkImageProcessorV1_ProcessImage(b *testing.B) {
	processor := NewImageProcessorV1(logger)

	for i := 0; i < b.N; i++ {
		processor.ProcessImage(sample)
	}
}

func TestImageProcessorV1_ProcessImage(t *testing.T) {
	if _, err := os.Stat(sample); errors.Is(err, os.ErrNotExist) {
		logger.Error("failed to open file")
	} else {
		type fields struct {
			logger *slog.Logger
		}
		type args struct {
			path string
		}

		field := fields{logger: logger}
		arg := args{path: sample}

		expected := [][]int{
			{0, 0, 0, 0, 4, 6, 0, 0, 0},
			{3, 0, 0, 0, 0, 0, 0, 8, 0},
			{0, 0, 0, 0, 7, 0, 0, 0, 0},
			{2, 0, 0, 0, 0, 0, 6, 0, 5},
			{0, 5, 0, 8, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 7, 0, 0},
			{0, 9, 7, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 5, 0, 0, 0, 3, 0},
			{4, 0, 6, 0, 0, 0, 0, 0, 0},
		}

		tests := []struct {
			name   string
			fields fields
			args   args
			want   [][]int
		}{
			{
				name:   "test_1",
				fields: field,
				args:   arg,
				want:   expected,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				processor := &ImageProcessorV1{
					logger: tt.fields.logger,
				}
				if got := processor.ProcessImage(tt.args.path); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ProcessImage() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
