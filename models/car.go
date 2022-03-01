package models

type Car struct {
	ID                int    `json:"cid"`
	Name              string `json:"name"`
	Brand             string `json:"brand"`
	FuelType          string `json:"fuelType"`
	YearOfManufacture int    `json:"yearOfManufacture"`
}
