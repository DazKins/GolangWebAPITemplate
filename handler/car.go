package handler

import (
	"CarAPI/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type carFetcher interface {
	FetchCar(carId string) (*model.Car, error)
}

type carHandler struct {
	carFetcher carFetcher
}

func NewCarHandler(carFetcher carFetcher) carHandler {
	return carHandler{
		carFetcher: carFetcher,
	}
}

func (api carHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	if id == "" {
		w.WriteHeader(400)
		w.Write([]byte("Car ID was not provided!"))
	}

	car, err := api.carFetcher.FetchCar(id)

	if err != nil {
	}

	body, _ := json.Marshal(car.ToDto())

	w.WriteHeader(200)
	w.Write(body)
}
