package model

import "github.com/google/uuid"

type CarId uuid.UUID

type CarIdDto string

func (c CarId) ToDto() CarIdDto {
	return CarIdDto(uuid.UUID(c).String())
}

type Manafacturer string

type ManafacturerDto string

func (m Manafacturer) ToDto() ManafacturerDto {
	return ManafacturerDto(string(m))
}

type Colour string

type ColourDto string

func (c Colour) ToDto() ColourDto {
	return ColourDto(string(c))
}

type Car struct {
	Id           CarId
	Manafacturer Manafacturer
	Colour       Colour
}

type CarDto struct {
	Id           CarIdDto        `json:"id"`
	Manafacturer ManafacturerDto `json:"manafacturer"`
	Colour       ColourDto       `json:"colour"`
}

func (c Car) ToDto() CarDto {
	return CarDto{
		Id:           c.Id.ToDto(),
		Manafacturer: c.Manafacturer.ToDto(),
		Colour:       c.Colour.ToDto(),
	}
}
