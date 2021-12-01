package service

import (
	"CarAPI/config"
	"CarAPI/model"
	"errors"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type dvlaService struct {
	config config.Config
	client *resty.Client
}

func NewDvlaService(config config.Config, client *resty.Client) dvlaService {
	return dvlaService{
		config: config,
		client: client,
	}
}

func (service dvlaService) FetchCar(carId string) (*model.Car, error) {
	getCarEndpoint := fmt.Sprintf("%s/car/%s", service.config.DvlaApiUrl, carId)

	type ResponseModel struct {
		CarId        string `json:"carId"`
		Manafacturer string `json:"manafacturer"`
		Color        string `json:"color"`
	}

	resp, err := service.client.R().
		SetResult(&ResponseModel{}).
		Get(getCarEndpoint)

	if err != nil {
		log.Printf("Error calling %s, error: %s", getCarEndpoint, err.Error())
		return nil, errors.New("Failed to fetch car")
	}

	car := resp.Result().(*ResponseModel)

	return &model.Car{
		Id:           car.CarId,
		Manafacturer: car.Manafacturer,
		Colour:       car.Color,
	}, nil
}
