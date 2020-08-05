package product

import "time"

// Product structure entity
type Product struct {
	ID           string
	Name         string
	Description  string
	Price        float64
	Images       []string
	CurrencyCode string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
