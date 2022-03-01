package car

import (
	"database/sql"
	"github.com/krogertechnology/krogo/pkg/errors"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"krogoCRUD2/models"
)

type Store struct{}

func New() *Store {
	return &Store{}
}

func (c Store) Get(ctx *krogo.Context) (*[]models.Car, error) {
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = ctx.DB().Query("SELECT * FROM car")

	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	var cars []models.Car

	for rows.Next() {
		var car models.Car
		err = rows.Scan(&car.ID, &car.Name, &car.Brand, &car.FuelType, &car.YearOfManufacture)
		if err != nil {
			return nil, errors.DB{Err: err}
		}
		cars = append(cars, car)
	}
	return &cars, nil
}

func (c Store) GetByID(ctx *krogo.Context, id int) (*models.Car, error) {
	var car models.Car
	err := ctx.DB().QueryRow("SELECT * FROM car where id=?", id).Scan(&car.ID, &car.Name, &car.Brand, &car.FuelType, &car.YearOfManufacture)

	if err == sql.ErrNoRows {
		return nil, errors.DB{Err: err}
	}
	return &car, nil
}

func (c Store) Create(ctx *krogo.Context, car models.Car) (*models.Car, error) {
	_, err := ctx.DB().Exec("INSERT INTO car (id, name, brand, fueltype, yearofmanufacture) VALUES(?,?,?,?,?)", car.ID, car.Name, car.Brand, car.FuelType, car.YearOfManufacture)
	if err != nil {
		return nil, errors.DB{Err: err}
	}
	return c.GetByID(ctx, car.ID)
}

func (c Store) Delete(ctx *krogo.Context, id int) error {
	_, err := ctx.DB().Exec("DELETE FROM car WHERE id=?", id)
	if err != nil {
		return errors.DB{Err: err}
	}
	return nil
}
