package main

import (
	"CarAPI/config"
	"CarAPI/handler"
	"CarAPI/service"
	"fmt"

	"github.com/go-resty/resty/v2"
	"go.uber.org/dig"
)

func main() {
	c := dig.New()

	c.Provide(handler.NewCarHandler)
	c.Provide(service.NewDvlaService)
	c.Provide(NewServer)
	c.Provide(config.GenerateConfig)
	c.Provide(resty.New)

	err := c.Invoke(func(api CarAPI) {
		api.Run()
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}
