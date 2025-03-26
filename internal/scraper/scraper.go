package scraper

import (
	"github.com/tedjuang/go-scrapy/internal/models"
)

// Scraper defines the interface for product scrapers
type Scraper interface {
	// ScrapeProduct scrapes a product from a URL
	ScrapeProduct(url string) (*models.Product, error)

	// ScrapeSearch scrapes search results for a keyword
	ScrapeSearch(keyword string, maxProducts int) ([]*models.Product, error)
}

// ScraperFactory creates a new scraper for a given website
type ScraperFactory struct {
	scrapers map[string]Scraper
}

// NewScraperFactory creates a new scraper factory with all supported scrapers
func NewScraperFactory() *ScraperFactory {
	factory := &ScraperFactory{
		scrapers: make(map[string]Scraper),
	}

	// Register all supported scrapers
	factory.scrapers["rakuten"] = NewRakutenScraper()

	return factory
}

// GetScraper returns a scraper for a given website
func (sf *ScraperFactory) GetScraper(website string) (Scraper, bool) {
	scraper, exists := sf.scrapers[website]
	return scraper, exists
}

// GetAllScrapers returns all registered scrapers
func (sf *ScraperFactory) GetAllScrapers() map[string]Scraper {
	return sf.scrapers
}
