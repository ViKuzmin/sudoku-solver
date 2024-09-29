package api

import (
	"github.com/gorilla/mux"
	"sudoku-solver/internal/handlers"
)

func CreateRoutes(imageHandler *handlers.ImageHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/shellscript", imageHandler.GetAndroidShellScript).Methods("POST")
	r.HandleFunc("/rawdata", imageHandler.GetRawAnswerData).Methods("POST")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()
	return r
}
