package user

import (
	"errors"
	"github.com/NegarMov/shopping-api/internal/model"
)

var (
	ErrDuplicatedUsername	= errors.New("the chosen username already exists")
	ErrWrongCredentials		= errors.New("incorrect username and/or password")
)

type User interface {
	SignUp(user model.User) (model.User, error)
	SignIn(username string, password string) (model.User, error)
}