package model

import (
	"time"
	"github.com/jackc/pgtype"
)

type State string
const (
	Completed State = "COMPLETED"
	Pending   State = "PENDING"
)

type Basket struct {
	ID      	uint			`json:"id"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	time.Time 		`json:"updated_at"`
	Data		pgtype.JSONB	`json:"data"`
	State		State			`json:"state"`
	UserID		uint			`json:"user_id"`
}
