package request

import (
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/NegarMov/shopping-api/internal/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type BasketCreate struct {
	Data	pgtype.JSONB	`json:"data"`
	State	model.State		`json:"state"`
}

func (r BasketCreate) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.Data, validation.Required, validation.Length(0, 2048)),
		validation.Field(&r.State, validation.Required, validation.In(model.Completed, model.Pending)),
	); err != nil {
		return fmt.Errorf("basket creation request validation failed - %w", err)
	}

	return nil
}

type BasketUpdate struct {
	Data	pgtype.JSONB	`json:"data"`
	State	model.State		`json:"state"`
}

func (r BasketUpdate) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.Data, validation.Length(0, 2048)),
		validation.Field(&r.State, validation.In(model.Completed, model.Pending)),
	); err != nil {
		return fmt.Errorf("basket update request validation failed - %w", err)
	}

	return nil
}