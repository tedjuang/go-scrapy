package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/tedjuang/go-scrapy/internal/models"
	"github.com/tedjuang/go-scrapy/internal/scraper"
	"github.com/tedjuang/go-scrapy/internal/storage"
)

func main() {
	// Define command line flags
	url := flag.String("url", "", "URL of the product to track")
	search := flag.String("search", "", "Search for products with this keyword")
	website := flag.String("website", "rakuten", "Website to scrape (e.g., rakuten)")
	maxResults := flag.Int("max", 10, "Maximum number of search results")
	dataDir := flag.String("data", "./data", "Directory to store data")

	// Parse command line flags
	flag.Parse()

	// Create data directory if it doesn't exist
	if err := os.MkdirAll(*dataDir, 0755); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	// Create a storage instance
	store, err := storage.NewJSONFileStorage(filepath.Join(*dataDir, "products.json"))
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	// Create a scraper factory
	factory := scraper.NewScraperFactory()

	// Get the appropriate scraper
	s, exists := factory.GetScraper(*website)
	if !exists {
		log.Fatalf("No scraper found for website: %s", *website)
	}

	// Process command based on flags
	if *url != "" {
		// Scrape a single product
		fmt.Printf("Scraping product from URL: %s\n", *url)
		product, err := s.ScrapeProduct(*url)
		if err != nil {
			log.Fatalf("Failed to scrape product: %v", err)
		}

		// Print product details
		printProduct(product)

		// Save to storage
		if err := store.Save(product); err != nil {
			log.Fatalf("Failed to save product: %v", err)
		}
		fmt.Println("Product saved successfully!")
	} else if *search != "" {
		// Search for products
		fmt.Printf("Searching for '%s' on %s (max: %d results)\n", *search, *website, *maxResults)
		products, err := s.ScrapeSearch(*search, *maxResults)
		if err != nil {
			log.Fatalf("Failed to search for products: %v", err)
		}

		// Print search results
		fmt.Printf("Found %d products:\n", len(products))
		for i, p := range products {
			fmt.Printf("\n--- Product %d ---\n", i+1)
			printProduct(p)

			// Save to storage
			if err := store.Save(p); err != nil {
				log.Printf("Warning: Failed to save product %s: %v", p.ID, err)
			}
		}
		fmt.Println("\nAll products saved successfully!")
	} else {
		// If no URL or search provided, show usage information
		flag.Usage()
		os.Exit(1)
	}
}

// printProduct displays product information
func printProduct(p *models.Product) {
	fmt.Printf("ID: %s\n", p.ID)
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("URL: %s\n", p.URL)
	fmt.Printf("Price: %.2f %s\n", p.CurrentPrice, p.Currency)
	fmt.Printf("Image URL: %s\n", p.ImageURL)
	if len(p.Description) > 100 {
		fmt.Printf("Description: %s...\n", p.Description[:100])
	} else {
		fmt.Printf("Description: %s\n", p.Description)
	}
	fmt.Printf("Last Updated: %s\n", p.LastUpdated.Format(time.RFC1123))
}
