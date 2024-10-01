package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWrapErrorWithStatus(t *testing.T) {
	type args struct {
		w          http.ResponseWriter
		err        error
		httpStatus int
	}

	w := httptest.NewRecorder()
	expectedResponse := ErrorResponse{
		Data:   "qwe",
		Result: "error",
	}

	var actualResponse ErrorResponse

	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_1",
			args: args{
				w:          w,
				err:        fmt.Errorf("qwe"),
				httpStatus: 502,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WrapErrorWithStatus(tt.args.w, tt.args.err, tt.args.httpStatus)
			json.Unmarshal(w.Body.Bytes(), &actualResponse)
			if actualResponse != expectedResponse || w.Code != tt.args.httpStatus {
				t.Errorf("Get response %v, want %v", actualResponse, expectedResponse)
			}
		})
	}
}

type ErrorResponse struct {
	Data   string `json:"data"`
	Result string `json:"result"`
}
