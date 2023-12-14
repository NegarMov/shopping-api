package basket

import (
	"errors"
	"github.com/NegarMov/shopping-api/internal/model"
)

var (
	ErrBasketNotFound	= errors.New("no basket with the given ID was found")
	ErrBasketCompleted	= errors.New("unable to update a basket marked as COMPLETED")
)

type Basket interface {
	GetAll() ([]model.Basket, error)
	Create(basket model.Basket) (model.Basket, error)
	Update(id uint, basket model.Basket) (model.Basket, error)
	Get(id uint) (model.Basket, error)
	Delete(id uint) (model.Basket, error)
}