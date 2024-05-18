package transaction

import (
	"time"
)

type (
	Transaction struct {
		ID          uint
		Description string
		RequestedAt time.Time
		CreatedAt   time.Time
	}
)

func New(description string, requestedAt time.Time) Transaction {
	return Transaction{
		Description: description,
		RequestedAt: requestedAt,
	}
}
