package invoiceheader

import "time"

// Model of invoiceheader
type Model struct {
	ID        int64     `json:"id"`
	Client    string    `json:"client"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}
