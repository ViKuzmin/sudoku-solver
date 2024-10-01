package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFound(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	w := httptest.NewRecorder()

	expectedResponse := ErrorResponse{
		Data:   "not found",
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
				w: w,
				r: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NotFound(tt.args.w, tt.args.r)

			json.Unmarshal(w.Body.Bytes(), &actualResponse)

			if w.Code != http.StatusNotFound || expectedResponse != actualResponse {
				t.Errorf("Get response %v, want %v", actualResponse, expectedResponse)
			}
		})
	}
}
