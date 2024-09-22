package api

import (
	"github.com/gorilla/mux"
	"ocr-test/internal/handlers"
)

func CreateRoutes(imageHandler *handlers.ImageProcessor) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/file", imageHandler.ProcessImage).Methods("POST")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler() //оборачиваем 404, для обработки NotFound
	return r
}
