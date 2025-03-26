package models

import (
	"time"
)

// Product represents an e-commerce product with price tracking
type Product struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	URL          string       `json:"url"`
	ImageURL     string       `json:"image_url"`
	Description  string       `json:"description"`
	PriceHistory []PricePoint `json:"price_history"`
	CurrentPrice float64      `json:"current_price"`
	Currency     string       `json:"currency"`
	LastUpdated  time.Time    `json:"last_updated"`
	Website      string       `json:"website"` // e.g., "rakuten"
}

// PricePoint represents a price at a specific point in time
type PricePoint struct {
	Price     float64   `json:"price"`
	Currency  string    `json:"currency"`
	Timestamp time.Time `json:"timestamp"`
}

// NewProduct creates a new product with default values
func NewProduct(id, name, url, website string, price float64, currency string) *Product {
	now := time.Now()
	return &Product{
		ID:           id,
		Name:         name,
		URL:          url,
		Website:      website,
		CurrentPrice: price,
		Currency:     currency,
		PriceHistory: []PricePoint{
			{
				Price:     price,
				Currency:  currency,
				Timestamp: now,
			},
		},
		LastUpdated: now,
	}
}

// UpdatePrice adds a new price point to the product's price history
func (p *Product) UpdatePrice(price float64, currency string) {
	now := time.Now()
	p.CurrentPrice = price
	p.Currency = currency
	p.LastUpdated = now

	// Add to price history
	p.PriceHistory = append(p.PriceHistory, PricePoint{
		Price:     price,
		Currency:  currency,
		Timestamp: now,
	})
}
