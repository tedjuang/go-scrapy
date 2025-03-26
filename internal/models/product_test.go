package models

import (
	"testing"
	"time"
)

func TestNewProduct(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		prodName string
		url      string
		website  string
		price    float64
		currency string
	}{
		{
			name:     "Valid product creation",
			id:       "test-123",
			prodName: "Test Product",
			url:      "https://example.com/product",
			website:  "rakuten",
			price:    99.99,
			currency: "JPY",
		},
		{
			name:     "Zero price product",
			id:       "test-456",
			prodName: "Free Product",
			url:      "https://example.com/free",
			website:  "rakuten",
			price:    0.0,
			currency: "JPY",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product := NewProduct(tt.id, tt.prodName, tt.url, tt.website, tt.price, tt.currency)

			// Validate product fields
			if product.ID != tt.id {
				t.Errorf("Expected ID %s, got %s", tt.id, product.ID)
			}
			if product.Name != tt.prodName {
				t.Errorf("Expected Name %s, got %s", tt.prodName, product.Name)
			}
			if product.URL != tt.url {
				t.Errorf("Expected URL %s, got %s", tt.url, product.URL)
			}
			if product.Website != tt.website {
				t.Errorf("Expected Website %s, got %s", tt.website, product.Website)
			}
			if product.CurrentPrice != tt.price {
				t.Errorf("Expected CurrentPrice %f, got %f", tt.price, product.CurrentPrice)
			}
			if product.Currency != tt.currency {
				t.Errorf("Expected Currency %s, got %s", tt.currency, product.Currency)
			}

			// Validate price history
			if len(product.PriceHistory) != 1 {
				t.Errorf("Expected 1 price history entry, got %d", len(product.PriceHistory))
			} else {
				if product.PriceHistory[0].Price != tt.price {
					t.Errorf("Expected price history price %f, got %f", tt.price, product.PriceHistory[0].Price)
				}
				if product.PriceHistory[0].Currency != tt.currency {
					t.Errorf("Expected price history currency %s, got %s", tt.currency, product.PriceHistory[0].Currency)
				}
			}
		})
	}
}

func TestUpdatePrice(t *testing.T) {
	// Create a product
	product := NewProduct("test-123", "Test Product", "https://example.com/product", "rakuten", 99.99, "JPY")

	// Initial state validation
	if len(product.PriceHistory) != 1 {
		t.Fatalf("Expected 1 price history entry, got %d", len(product.PriceHistory))
	}

	// Record the initial timestamp
	initialTimestamp := product.LastUpdated

	// Wait a short time to ensure timestamp changes
	time.Sleep(10 * time.Millisecond)

	// Update the price
	newPrice := 89.99
	newCurrency := "JPY"
	product.UpdatePrice(newPrice, newCurrency)

	// Validate updated fields
	if product.CurrentPrice != newPrice {
		t.Errorf("Expected CurrentPrice %f, got %f", newPrice, product.CurrentPrice)
	}

	if product.Currency != newCurrency {
		t.Errorf("Expected Currency %s, got %s", newCurrency, product.Currency)
	}

	// Validate that LastUpdated changed
	if !product.LastUpdated.After(initialTimestamp) {
		t.Errorf("Expected LastUpdated to be after initial timestamp")
	}

	// Validate price history
	if len(product.PriceHistory) != 2 {
		t.Errorf("Expected 2 price history entries, got %d", len(product.PriceHistory))
	} else {
		lastEntry := product.PriceHistory[len(product.PriceHistory)-1]
		if lastEntry.Price != newPrice {
			t.Errorf("Expected last price history price %f, got %f", newPrice, lastEntry.Price)
		}
		if lastEntry.Currency != newCurrency {
			t.Errorf("Expected last price history currency %s, got %s", newCurrency, lastEntry.Currency)
		}
	}
}
