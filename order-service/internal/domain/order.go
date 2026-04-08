package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	Pending   OrderStatus = "pending"
	Confirmed OrderStatus = "confirmed"
	Cancelled OrderStatus = "cancelled"
)

type Order struct {
	ID          uuid.UUID   `db:"id"`
	ItemName    string      `db:"item_name"`
	Description string      `db:"description"`
	TotalAmount float64     `db:"total_amount"`
	Status      OrderStatus `db:"status"` // pending, confirmed, cancelled
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
}
