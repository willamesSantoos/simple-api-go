package routes

import (
	"github.com/gorilla/mux"
	"github.com/willamesSantoos/simple-api-go/internal/handlers"
)

func InitializeRoutes(r *mux.Router) {
	r.HandleFunc("/cars", handlers.GetCars)
	r.HandleFunc("/cars/{id:[0-9]+}", handlers.GetCar)
}