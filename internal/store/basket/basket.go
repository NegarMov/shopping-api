package basket

import (
	"errors"
	"github.com/NegarMov/shopping-api/internal/model"
)

var (
	ErrBasketNotFound	= errors.New("no basket with the given ID was found")
	ErrBasketCompleted	= errors.New("unable to update a basket marked as COMPLETED")
	ErrAccessDenied		= errors.New("you don't have access to this basket!")
)

type Basket interface {
	GetAll(user_id uint) ([]model.Basket, error)
	Create(basket model.Basket) (model.Basket, error)
	Update(id uint, basket model.Basket) (model.Basket, error)
	Get(id uint, user_id uint) (model.Basket, error)
	Delete(id uint, user_id uint) (model.Basket, error)
}