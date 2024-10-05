package delivery

import "github.com/Inspirate789/ds-lab2/internal/models"

type CarDTO struct {
	ID                 int64          `json:"id"`
	CarUID             string         `json:"car_uid"`
	Brand              string         `json:"brand"`
	Model              string         `json:"model"`
	RegistrationNumber string         `json:"registration_number"`
	Power              uint64         `json:"power"`
	Price              uint64         `json:"price"`
	Type               models.CarType `json:"type"`
	Availability       bool           `json:"availability"`
}

func NewCarDTO(car models.Car) CarDTO {
	return CarDTO{
		ID:                 car.ID,
		CarUID:             car.CarUID,
		Brand:              car.Brand,
		Model:              car.Model,
		RegistrationNumber: car.RegistrationNumber,
		Power:              car.Power,
		Price:              car.Price,
		Type:               car.Type,
		Availability:       car.Availability,
	}
}

type CarsDTO struct {
	Items []models.Car `json:"items"`
	Count uint64       `json:"count"`
}

func NewCarsDTO(cars []models.Car, totalCount uint64) CarsDTO {
	return CarsDTO{
		Items: cars,
		Count: totalCount,
	}
}
