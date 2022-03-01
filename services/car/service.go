package car

import (
	"github.com/krogertechnology/krogo/pkg/krogo"
	"krogoCRUD2/models"
	"krogoCRUD2/stores"
)

type Service struct {
	datastoreCar stores.Car
}

func New(datastoreCar stores.Car) Service {
	return Service{datastoreCar: datastoreCar}
}

func (c Service) Get(ctx *krogo.Context) (*[]models.Car, error) {
	cars, err := c.datastoreCar.Get(ctx)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (c Service) GetById(ctx *krogo.Context, id int) (*models.Car, error) {
	car, err := c.datastoreCar.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (c Service) Insert(ctx *krogo.Context, car models.Car) (*models.Car, error) {
	resp, err := c.datastoreCar.Create(ctx, car)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c Service) Delete(ctx *krogo.Context, id int) error {
	err := c.datastoreCar.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
