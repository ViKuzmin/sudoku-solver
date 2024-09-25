package image_processor

import (
	"log/slog"
	"os"
	"testing"
)

func BenchmarkImageProcessorV1_ProcessImage(b *testing.B) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	processor := NewImageProcessorV1(logger)

	for i := 0; i < b.N; i++ {
		processor.ProcessImage("sample.jpg")
	}
}
