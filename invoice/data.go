package invoice

import (
	"time"

	"github.com/google/uuid"
)

const (
	PENDING = byte(iota)
	PAID
	CANCELD
)

type Invoice struct {
	id          uuid.UUID
	number      string
	date        time.Time
	status      int
	totalAmount float64
}

// Getter methods

func (i *Invoice) GetID() uuid.UUID {
	return i.id
}

func (i *Invoice) GetNumber() string {
	return i.number
}

func (i *Invoice) GetDate() time.Time {
	return i.date
}

func (i *Invoice) GetStatus() int {
	return i.status
}

func (i *Invoice) GetTotalAmount() float64 {
	return i.totalAmount
}

// Setter methods

func (i *Invoice) SetID(id uuid.UUID) {
	i.id = id
}

func (i *Invoice) SetNumber(number string) {
	i.number = number
}

func (i *Invoice) SetDate(date time.Time) {
	i.date = date
}

func (i *Invoice) SetStatus(status int) {
	i.status = status
}

func (i *Invoice) SetTotalAmount(totalAmount float64) {
	i.totalAmount = totalAmount
}
