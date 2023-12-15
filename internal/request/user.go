package request

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Credentials struct {
	Username	string	`json:"username"`
	Password    string	`json:"password"`
}

func (u Credentials) Validate() error {
	if err := validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required, validation.Length(1, 100)),
		validation.Field(&u.Password, validation.Required, validation.Length(1, 100)),
	); err != nil {
		return fmt.Errorf("user creation request validation failed - %w", err)
	}

	return nil
}