package model

import (
	"time"
)

type State string
const (
	Completed State = "COMPLETED"
	Pending   State = "PENDING"
)

type Basket struct {
	ID      	uint		`json:"id"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time 	`json:"updated_at"`
	Data		string		`json:"data"`
	State		State		`json:"state"`
}
