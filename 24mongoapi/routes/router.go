package routes

import (
	"github.com/dusariudaykumar/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/movies", controller.GetAllMoviesController).Methods("GET")
	r.HandleFunc("/movies/{id}", controller.GetMovie).Methods("GET")
	r.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controller.MarkMovieAsWatched).Methods("PUT")
	r.HandleFunc("/movies/{id}", controller.DeleteOneMovie).Methods("DELETE")

	return r
}
