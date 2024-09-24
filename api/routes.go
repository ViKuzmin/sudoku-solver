package api

import (
	"github.com/gorilla/mux"
	"ocr-test/internal/handlers"
)

func CreateRoutes(imageHandler *handlers.ImageHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/file", imageHandler.ProcessImage).Methods("POST")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()
	return r
}
