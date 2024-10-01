package handlers

import (
	"bytes"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"sudoku-solver/internal/image_processing/image_processor"
	"sudoku-solver/internal/solver/sudoku_solver"
	"testing"
)

var sample = "sample.jpg"
var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
var proc = image_processor.NewImageProcessorV1(logger)
var solver = sudoku_solver.NewSolver(logger)
var handler = ImageHandler{
	logger:    logger,
	Processor: proc,
	Solver:    solver,
}

func TestNewImageHandler(t *testing.T) {
	type args struct {
		logger    *slog.Logger
		processor *image_processor.ImageProcessorV1
	}

	tests := []struct {
		name string
		args args
		want *ImageHandler
	}{
		{
			name: "test_1",
			args: args{
				logger:    logger,
				processor: proc,
			},
			want: &handler,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewImageHandler(tt.args.logger, tt.args.processor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImageHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImageHandler_GetRawAnswerData(t *testing.T) {
	w := httptest.NewRecorder()
	r := getBody(t)

	tests := []struct {
		name    string
		handler ImageHandler
		args    args
		wanted  wanted
	}{
		{
			name:    "test_1",
			handler: handler,
			args: args{
				w: w,
				r: r,
			},
			wanted: wanted{
				// TODO correct answer
				//data: "[[8,7,1,9,4,6,2,5,3],[3,6,4,1,2,5,9,8,7],[9,2,5,3,7,8,1,6,4],[2,4,8,7,1,3,6,9,5],[7,5,9,8,6,2,3,4,1],[6,1,3,4,5,9,7,2,8],[5,9,7,6,3,4,8,1,2],[1,8,2,5,9,7,4,3,6],[4,3,6,2,8,1,5,7,9]]",
				data: "[[1,2,3,4,5,6,7,8,9],[4,5,6,7,8,9,1,2,3],[7,8,9,1,2,3,4,5,6],[2,1,4,3,6,5,8,9,7],[3,6,5,8,9,7,2,1,4],[8,9,7,2,1,4,3,6,5],[5,3,1,6,4,2,9,7,8],[6,4,2,9,7,8,5,3,1],[9,7,8,5,3,1,6,4,2]]",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.handler.GetRawAnswerData(tt.args.w, tt.args.r)
			if actual := w.Body.String(); actual != tt.wanted.data {
				t.Errorf("Raw data = %v, want %v", actual, tt.wanted.data)
			}
		})
	}
}

func TestImageHandler_GetAndroidShellScript(t *testing.T) {
	w := httptest.NewRecorder()
	r := getBody(t)

	tests := []struct {
		name    string
		handler ImageHandler
		args    args
		wanted  wanted
	}{
		{
			name:    "test_1",
			handler: handler,
			args: args{
				w: w,
				r: r,
			},
			wanted: wanted{
				data: scriptExpected,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.handler.GetAndroidShellScript(tt.args.w, tt.args.r)

			if actual := w.Body.String(); actual != tt.wanted.data {
				t.Errorf("Script = %v, want %v", actual, tt.wanted.data)
			}
		})
	}
}

func getBody(t *testing.T) *http.Request {
	file, err := os.Open(sample)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", sample)
	if err != nil {
		t.Fatal(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/rawdata", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req
}

type args struct {
	w http.ResponseWriter
	r *http.Request
}

type wanted struct {
	data string
}

// TODO correct answer
// var scriptExpected = "input tap 90 505; input tap 889 2011; input tap 202 505; input tap 773 2011; input tap 314 505; input tap 77 2011; input tap 426 505; input tap 1005 2011; input tap 538 505; input tap 425 2011; input tap 650 505; input tap 657 2011; input tap 762 505; input tap 193 2011; input tap 874 505; input tap 541 2011; input tap 986 505; input tap 309 2011; input tap 130 619; input tap 309 2011; input tap 242 619; input tap 657 2011; input tap 354 619; input tap 425 2011; input tap 466 619; input tap 77 2011; input tap 578 619; input tap 193 2011; input tap 690 619; input tap 541 2011; input tap 802 619; input tap 1005 2011; input tap 914 619; input tap 889 2011; input tap 1026 619; input tap 773 2011; input tap 130 733; input tap 1005 2011; input tap 242 733; input tap 193 2011; input tap 354 733; input tap 541 2011; input tap 466 733; input tap 309 2011; input tap 578 733; input tap 773 2011; input tap 690 733; input tap 889 2011; input tap 802 733; input tap 77 2011; input tap 914 733; input tap 657 2011; input tap 1026 733; input tap 425 2011; input tap 130 847; input tap 193 2011; input tap 242 847; input tap 425 2011; input tap 354 847; input tap 889 2011; input tap 466 847; input tap 773 2011; input tap 578 847; input tap 77 2011; input tap 690 847; input tap 309 2011; input tap 802 847; input tap 657 2011; input tap 914 847; input tap 1005 2011; input tap 1026 847; input tap 541 2011; input tap 130 961; input tap 773 2011; input tap 242 961; input tap 541 2011; input tap 354 961; input tap 1005 2011; input tap 466 961; input tap 889 2011; input tap 578 961; input tap 657 2011; input tap 690 961; input tap 193 2011; input tap 802 961; input tap 309 2011; input tap 914 961; input tap 425 2011; input tap 1026 961; input tap 77 2011; input tap 130 1075; input tap 657 2011; input tap 242 1075; input tap 77 2011; input tap 354 1075; input tap 309 2011; input tap 466 1075; input tap 425 2011; input tap 578 1075; input tap 541 2011; input tap 690 1075; input tap 1005 2011; input tap 802 1075; input tap 773 2011; input tap 914 1075; input tap 193 2011; input tap 1026 1075; input tap 889 2011; input tap 130 1189; input tap 541 2011; input tap 242 1189; input tap 1005 2011; input tap 354 1189; input tap 773 2011; input tap 466 1189; input tap 657 2011; input tap 578 1189; input tap 309 2011; input tap 690 1189; input tap 425 2011; input tap 802 1189; input tap 889 2011; input tap 914 1189; input tap 77 2011; input tap 1026 1189; input tap 193 2011; input tap 130 1303; input tap 77 2011; input tap 242 1303; input tap 889 2011; input tap 354 1303; input tap 193 2011; input tap 466 1303; input tap 541 2011; input tap 578 1303; input tap 1005 2011; input tap 690 1303; input tap 773 2011; input tap 802 1303; input tap 425 2011; input tap 914 1303; input tap 309 2011; input tap 1026 1303; input tap 657 2011; input tap 130 1417; input tap 425 2011; input tap 242 1417; input tap 309 2011; input tap 354 1417; input tap 657 2011; input tap 466 1417; input tap 193 2011; input tap 578 1417; input tap 889 2011; input tap 690 1417; input tap 77 2011; input tap 802 1417; input tap 541 2011; input tap 914 1417; input tap 773 2011; input tap 1026 1417; input tap 1005 2011; "
var scriptExpected = "input tap 90 505; input tap 77 2011; input tap 202 505; input tap 193 2011; input tap 314 505; input tap 309 2011; input tap 426 505; input tap 425 2011; input tap 538 505; input tap 541 2011; input tap 650 505; input tap 657 2011; input tap 762 505; input tap 773 2011; input tap 874 505; input tap 889 2011; input tap 986 505; input tap 1005 2011; input tap 130 619; input tap 425 2011; input tap 242 619; input tap 541 2011; input tap 354 619; input tap 657 2011; input tap 466 619; input tap 773 2011; input tap 578 619; input tap 889 2011; input tap 690 619; input tap 1005 2011; input tap 802 619; input tap 77 2011; input tap 914 619; input tap 193 2011; input tap 1026 619; input tap 309 2011; input tap 130 733; input tap 773 2011; input tap 242 733; input tap 889 2011; input tap 354 733; input tap 1005 2011; input tap 466 733; input tap 77 2011; input tap 578 733; input tap 193 2011; input tap 690 733; input tap 309 2011; input tap 802 733; input tap 425 2011; input tap 914 733; input tap 541 2011; input tap 1026 733; input tap 657 2011; input tap 130 847; input tap 193 2011; input tap 242 847; input tap 77 2011; input tap 354 847; input tap 425 2011; input tap 466 847; input tap 309 2011; input tap 578 847; input tap 657 2011; input tap 690 847; input tap 541 2011; input tap 802 847; input tap 889 2011; input tap 914 847; input tap 1005 2011; input tap 1026 847; input tap 773 2011; input tap 130 961; input tap 309 2011; input tap 242 961; input tap 657 2011; input tap 354 961; input tap 541 2011; input tap 466 961; input tap 889 2011; input tap 578 961; input tap 1005 2011; input tap 690 961; input tap 773 2011; input tap 802 961; input tap 193 2011; input tap 914 961; input tap 77 2011; input tap 1026 961; input tap 425 2011; input tap 130 1075; input tap 889 2011; input tap 242 1075; input tap 1005 2011; input tap 354 1075; input tap 773 2011; input tap 466 1075; input tap 193 2011; input tap 578 1075; input tap 77 2011; input tap 690 1075; input tap 425 2011; input tap 802 1075; input tap 309 2011; input tap 914 1075; input tap 657 2011; input tap 1026 1075; input tap 541 2011; input tap 130 1189; input tap 541 2011; input tap 242 1189; input tap 309 2011; input tap 354 1189; input tap 77 2011; input tap 466 1189; input tap 657 2011; input tap 578 1189; input tap 425 2011; input tap 690 1189; input tap 193 2011; input tap 802 1189; input tap 1005 2011; input tap 914 1189; input tap 773 2011; input tap 1026 1189; input tap 889 2011; input tap 130 1303; input tap 657 2011; input tap 242 1303; input tap 425 2011; input tap 354 1303; input tap 193 2011; input tap 466 1303; input tap 1005 2011; input tap 578 1303; input tap 773 2011; input tap 690 1303; input tap 889 2011; input tap 802 1303; input tap 541 2011; input tap 914 1303; input tap 309 2011; input tap 1026 1303; input tap 77 2011; input tap 130 1417; input tap 1005 2011; input tap 242 1417; input tap 773 2011; input tap 354 1417; input tap 889 2011; input tap 466 1417; input tap 541 2011; input tap 578 1417; input tap 309 2011; input tap 690 1417; input tap 77 2011; input tap 802 1417; input tap 657 2011; input tap 914 1417; input tap 425 2011; input tap 1026 1417; input tap 193 2011; "
