package scraper

import (
	"testing"
)

func TestNewScraperFactory(t *testing.T) {
	factory := NewScraperFactory()

	if factory == nil {
		t.Fatal("Expected factory to be non-nil")
	}

	if factory.scrapers == nil {
		t.Fatal("Expected scrapers map to be initialized")
	}

	// Check if Rakuten scraper is registered
	if _, exists := factory.scrapers["rakuten"]; !exists {
		t.Error("Expected Rakuten scraper to be registered")
	}
}

func TestGetScraper(t *testing.T) {
	factory := NewScraperFactory()

	tests := []struct {
		name         string
		website      string
		expectExists bool
	}{
		{
			name:         "Get Rakuten scraper",
			website:      "rakuten",
			expectExists: true,
		},
		{
			name:         "Get non-existent scraper",
			website:      "nonexistent",
			expectExists: false,
		},
		{
			name:         "Empty website name",
			website:      "",
			expectExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraper, exists := factory.GetScraper(tt.website)

			if exists != tt.expectExists {
				t.Errorf("Expected exists=%v, got %v", tt.expectExists, exists)
			}

			if tt.expectExists && scraper == nil {
				t.Error("Expected scraper to be non-nil when it exists")
			}

			if !tt.expectExists && scraper != nil {
				t.Error("Expected scraper to be nil when it doesn't exist")
			}
		})
	}
}

func TestGetAllScrapers(t *testing.T) {
	factory := NewScraperFactory()

	scrapers := factory.GetAllScrapers()

	if scrapers == nil {
		t.Fatal("Expected scrapers map to be non-nil")
	}

	// There should be at least the Rakuten scraper
	if len(scrapers) < 1 {
		t.Errorf("Expected at least 1 scraper, got %d", len(scrapers))
	}

	// Check if Rakuten scraper is in the map
	if _, exists := scrapers["rakuten"]; !exists {
		t.Error("Expected Rakuten scraper to be in the map")
	}
}
