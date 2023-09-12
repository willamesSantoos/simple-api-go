package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/willamesSantoos/simple-api-go/internal/repository"
)

func GetCars(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")

	jsonResponse, jsonError := json.Marshal(repository.GetListCar())

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

func GetCar(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(request)
	id := vars["id"]

	for _, car := range *repository.GetListCar() {
		if strconv.Itoa(car.Id) == id {
			jsonResponse, jsonError := json.Marshal(car)

			if jsonError != nil {
				fmt.Println("Unable to encode JSON")
			}

			response.WriteHeader(http.StatusOK)
			response.Write(jsonResponse)
		}
	}
}
