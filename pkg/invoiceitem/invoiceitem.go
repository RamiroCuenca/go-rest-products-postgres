package invoiceitem

import "time"

// Model of invoiceitem
type Model struct {
	ID              int64     `json:"id"`
	InvoiceHeaderID int64     `json:"invoice_header_id"`
	ProductID       int64     `json:"product_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdateAt        time.Time `json:"updated_at"`
}
