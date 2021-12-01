package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type carHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type carAPI struct {
	carHandler carHandler
}

func NewCarAPI(carHandler carHandler) carAPI {
	return carAPI{
		carHandler: carHandler,
	}
}

func (api carAPI) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/car/{id}", api.carHandler.Get)

	http.ListenAndServe(":3000", r)
}
