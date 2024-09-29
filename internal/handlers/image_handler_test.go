package handlers

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
var sample = "sample.jpg"

func TestImageHandler_GetRawSolvedData(t *testing.T) {

	//pr, pw := io.Pipe()
	//
	//writer := multipart.NewWriter(pw)
	//
	//go func() {
	//	defer writer.Close()
	//	// We create the form data field 'file'
	//	// which returns another writer to write the actual file
	//	part, err := writer.CreateFormFile("file", "sample.png")
	//	if err != nil {
	//		t.Error(err)
	//	}
	//
	//	// https://yourbasic.org/golang/create-image/
	//	img := image.Decode(part)
	//
	//	// Encode() takes an io.Writer.
	//	// We pass the multipart field
	//	// 'fileupload' that we defined
	//	// earlier which, in turn, writes
	//	// to our io.Pipe
	//	err = png.Encode(part, img)
	//	if err != nil {
	//		t.Error(err)
	//	}
	//}()
	//
	//request := httptest.NewRequest("POST", "/rawdata", pr)
	//request.Header.Add("Content-Type", writer.FormDataContentType())
	//
	//response := httptest.NewRecorder()
	//
	//processorV1 := image_processor.NewImageProcessorV1(logger)
	//solver := sudoku_solver.NewSolver(logger)
	//
	//processor := &ImageHandler{
	//	logger:    logger,
	//	Processor: processorV1,
	//	Solver:    solver,
	//}
	//
	//processor.GetRawAnswerData(response, request)
	//
	//t.Log("It should respond with an HTTP status code of 200")
	//if response.Code != 200 {
	//	//t.Errorf("Expected %s, received %d", 200, response.Code)
	//}
	//t.Log("It should create a file named 'someimg.png' in uploads folder")
	//if _, err := os.Stat("./uploads/someimg.png"); os.IsNotExist(err) {
	//	t.Error("Expected file ./uploads/someimg.png' to exist")
	//}

	//

	//

	//processorV1 := image_processor.NewImageProcessorV1(logger)
	//solver := sudoku_solver.NewSolver(logger)
	//
	//ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusOK)
	//}))
	//defer ts.Close()
	//
	//type fields struct {
	//	logger    *slog.Logger
	//	Processor *image_processor.ImageProcessorV1
	//	Solver    *sudoku_solver.Solver
	//}
	//type args struct {
	//	w http.ResponseWriter
	//	r *http.Request
	//}
	//tests := []struct {
	//	name   string
	//	fields fields
	//	args   args
	//}{
	//	{
	//		name: "test_1",
	//		fields: fields{
	//			logger:    logger,
	//			Processor: processorV1,
	//			Solver:    solver,
	//		},
	//		args: args{
	//			w: httptest.NewRecorder(),
	//			r: getRequest(),
	//		},
	//	},
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		processor := &ImageHandler{
	//			logger:    tt.fields.logger,
	//			Processor: tt.fields.Processor,
	//			Solver:    tt.fields.Solver,
	//		}
	//		processor.GetRawAnswerData(tt.args.w, tt.args.r)
	//	})
	//}
}

func getRequest() *http.Request {
	reqBody := []byte(`{"key":"value"}`)

	//request, _ := http.NewRequest("POST", "/rawdata", bytes.NewReader(reqBody))
	request := httptest.NewRequest("POST", "/rawdata", bytes.NewReader(reqBody))

	request.FormFile("file")

	return request
}
