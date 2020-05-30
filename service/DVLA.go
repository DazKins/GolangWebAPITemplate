package service

import (
	"CarAPI/config"
	"errors"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type DvlaService interface {
	GetCar(carId string) (*Car, error)
}

type dvlaService struct {
	Config config.Config
	Client *resty.Client
}

type Car struct {
	Id           string `json:"id"`
	Manafacturer string `json:"manafacturer"`
	Colour       string `json:"colour"`
}

func NewDvlaService(config config.Config, client *resty.Client) DvlaService {
	return dvlaService{
		Config: config,
		Client: client,
	}
}

func (service dvlaService) GetCar(carId string) (*Car, error) {
	getCarEndpoint := fmt.Sprintf("%s/car/%s", service.Config.DvlaApiUrl, carId)

	type ResponseModel struct {
		CarId        string `json:"carId"`
		Manafacturer string `json:"manafacturer"`
		Color        string `json:"color"`
	}

	resp, err := service.Client.R().
		SetResult(&ResponseModel{}).
		Get(getCarEndpoint)

	if err != nil {
		log.Printf("Error calling %s, error: %s", getCarEndpoint, err.Error())
		return nil, errors.New("Failed to fetch car")
	}

	car := resp.Result().(*ResponseModel)

	return &Car{
		Id:           car.CarId,
		Manafacturer: car.Manafacturer,
		Colour:       car.Color,
	}, nil
}
