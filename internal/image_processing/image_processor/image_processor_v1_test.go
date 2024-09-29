package image_processor

import (
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
		res := processor.ProcessImage(sample)
		if res == nil {
			break
		}
	}
}

func TestImageProcessorV1_ProcessImage(t *testing.T) {
	type fields struct {
		logger *slog.Logger
	}
	type args struct {
		path string
	}

	field := fields{logger: logger}

	//expected := [][]int{
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

	expected := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
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
			args: args{
				path: sample,
			},
			want: expected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := &ImageProcessorV1{
				logger: tt.fields.logger,
			}
			if got := processor.ProcessImage(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAndroidShellScript() = %v, want %v", got, tt.want)
			}
		})
	}
}
