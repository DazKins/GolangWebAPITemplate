package handler

import (
	"CarAPI/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CarHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type carHandler struct {
	DvlaService service.DvlaService
}

func NewCarHandler(dvlaService service.DvlaService) CarHandler {
	return carHandler{
		DvlaService: dvlaService,
	}
}

func (api carHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	if id == "" {
		w.WriteHeader(400)
		w.Write([]byte("Car ID was not provided!"))
	}

	car, err := api.DvlaService.GetCar(id)

	if err != nil {
	}

	body, _ := json.Marshal(car)

	w.WriteHeader(200)
	w.Write(body)
}
