package main

import (
	"CarAPI/handler"
	"net/http"

	"github.com/gorilla/mux"
)

type CarAPI struct {
	CarHandler handler.CarHandler
}

func NewServer(carHandler handler.CarHandler) CarAPI {
	return CarAPI{
		CarHandler: carHandler,
	}
}

func (api CarAPI) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/car/{id}", api.CarHandler.Get)

	http.ListenAndServe(":3000", r)
}
