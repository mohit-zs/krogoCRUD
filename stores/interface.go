package stores

import (
	"github.com/krogertechnology/krogo/pkg/krogo"
	"krogoCRUD2/models"
)

type Car interface {
	Get(ctx *krogo.Context) (*[]models.Car, error)
	GetByID(ctx *krogo.Context, id int) (*models.Car, error)
	Create(ctx *krogo.Context, car models.Car) (*models.Car, error)
	Delete(ctx *krogo.Context, id int) error
}
