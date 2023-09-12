package repository

import "github.com/willamesSantoos/simple-api-go/internal/models"

var listCar = []models.Car{
	{
        Id: 1,
		Fabricator: "fabricator 01",
		Model: "model 01",
		Year: 2022,
    },
	{
        Id: 2,
		Fabricator: "fabricator 02",
		Model: "model 02",
		Year: 2022,
    },
	{
        Id: 3,
		Fabricator: "fabricator 03",
		Model: "model 03",
		Year: 2022,
    },
}

func GetListCar () *[]models.Car {
	return &listCar
}