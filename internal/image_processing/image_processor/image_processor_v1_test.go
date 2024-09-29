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
		{
			name:   "test_2",
			fields: field,
			args: args{
				path: "",
			},
			want: nil,
		},
		{
			name:   "test_3",
			fields: field,
			args: args{
				path: "tt.txt",
			},
			want: nil,
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

func TestNewImageProcessorV1(t *testing.T) {
	type args struct {
		logger *slog.Logger
	}
	tests := []struct {
		name string
		args args
		want *ImageProcessorV1
	}{
		{
			name: "test_1",
			args: args{
				logger: logger,
			},
			want: NewImageProcessorV1(logger),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewImageProcessorV1(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImageProcessorV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
