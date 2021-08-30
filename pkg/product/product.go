package product

import "time"

// Model of product
type Model struct {
	ID           uint64    `json:"id"`
	Name         string    `json:"name"`
	Observations string    `json:"observations"`
	Price        int64     `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
