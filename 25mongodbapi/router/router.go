package router

import (
	"mongodbapi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("api/v1", controller.ServeHome).Methods("GET")
	router.HandleFunc("api/v1/movie", controller.CreateMovies).Methods("POST")
	router.HandleFunc("api/v1/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("api/v1/movie/{id}", controller.UpdateAMovie).Methods("PUT")
	router.HandleFunc("api/v1/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("api/v1/movies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
