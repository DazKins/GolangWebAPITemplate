package main

import (
	"CarAPI/config"
	"CarAPI/handler"
	"CarAPI/service"

	"github.com/go-resty/resty/v2"
)

func main() {
	config := config.GenerateConfig()
	resty := resty.New()
	dvlaService := service.NewDvlaService(config, resty)
	carHandler := handler.NewCarHandler(dvlaService)
	carApi := NewCarAPI(carHandler)

	carApi.Run()
}
