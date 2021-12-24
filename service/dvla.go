package service

import (
	"CarAPI/config"
	"CarAPI/model"
	"errors"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
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

func (service dvlaService) FetchCar(id string) (*model.Car, error) {
	getCarEndpoint := fmt.Sprintf("%s/car/%s", service.config.DvlaApiUrl, id)

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

	responseBody := resp.Result().(*ResponseModel)

	carId, err := uuid.Parse(responseBody.CarId)

	if err != nil {
		return nil, fmt.Errorf("Error parsing response car ID: %w", err)
	}

	return &model.Car{
		Id:           model.CarId(carId),
		Manafacturer: model.Manafacturer(responseBody.Manafacturer),
		Colour:       model.Colour(responseBody.Color),
	}, nil
}
