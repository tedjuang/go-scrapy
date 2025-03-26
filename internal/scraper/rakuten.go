package scraper

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/tedjuang/go-scrapy/internal/models"
)

// RakutenScraper implements scraper for Rakuten JP
type RakutenScraper struct {
	collector *colly.Collector
}

// NewRakutenScraper creates a new instance of RakutenScraper
func NewRakutenScraper() *RakutenScraper {
	// Create a new collector with custom settings
	c := colly.NewCollector(
		// Update: Allow more domains including search domain and books domain
		colly.AllowedDomains("www.rakuten.co.jp", "item.rakuten.co.jp", "search.rakuten.co.jp", "books.rakuten.co.jp"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"),
		colly.MaxDepth(2),
		// Allow redirects to other rakuten subdomains
		colly.AllowURLRevisit(),
	)

	// Set rate limiting to be respectful
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*rakuten.*",
		Parallelism: 2,
		Delay:       2 * time.Second,
	})

	return &RakutenScraper{
		collector: c,
	}
}

// ScrapeProduct scrapes a product from Rakuten JP based on its URL
func (rs *RakutenScraper) ScrapeProduct(url string) (*models.Product, error) {
	var product *models.Product
	var err error

	// Update: Use more selectors for product name to handle different page structures
	rs.collector.OnHTML("h1.item-name, h1#item-name, h1[itemprop='name'], span.item-name, h1.booksTitle", func(e *colly.HTMLElement) {
		productName := strings.TrimSpace(e.Text)
		// Create a temporary product, we'll populate it with more data as we scrape
		id := extractProductID(url)
		product = models.NewProduct(id, productName, url, "rakuten", 0, "JPY")
	})

	// Update: More price selectors to handle different page structures
	rs.collector.OnHTML(".price-box, #priceCalculationConfig, span[itemprop='price'], .price, .itemPrice, #priceAmount", func(e *colly.HTMLElement) {
		priceText := e.ChildText(".price")
		if priceText == "" {
			priceText = e.ChildText(".price-value")
		}
		if priceText == "" {
			priceText = e.Text
		}

		// Check if there's a data attribute for price
		if priceText == "" && e.Attr("data-price") != "" {
			priceText = e.Attr("data-price")
		}

		// Clean up the price text and extract just the numbers
		price := extractPrice(priceText)
		if price > 0 && product != nil {
			product.UpdatePrice(price, "JPY")
		}
	})

	// Update: More image selectors
	rs.collector.OnHTML("meta[property='og:image'], img.rakuten-main-product-image, img#imageURL", func(e *colly.HTMLElement) {
		if product != nil {
			imageURL := e.Attr("content")
			if imageURL == "" {
				imageURL = e.Attr("src")
			}
			product.ImageURL = imageURL
		}
	})

	// Update: More description selectors
	rs.collector.OnHTML("#item-description, .item-description, .item-details, .item-info, [itemprop='description'], #itemCaption", func(e *colly.HTMLElement) {
		if product != nil {
			doc := goquery.NewDocumentFromNode(e.DOM.Nodes[0])
			product.Description = strings.TrimSpace(doc.Text())
		}
	})

	rs.collector.OnError(func(r *colly.Response, err error) {
		log.Printf("Error scraping %s: %v", r.Request.URL, err)
	})

	// Start the scraping
	rs.collector.Visit(url)
	rs.collector.Wait()

	if product == nil {
		return nil, fmt.Errorf("failed to scrape product from URL: %s", url)
	}

	return product, err
}

// ScrapeSearch scrapes search results from Rakuten
func (rs *RakutenScraper) ScrapeSearch(keyword string, maxProducts int) ([]*models.Product, error) {
	var products []*models.Product
	count := 0

	searchURL := fmt.Sprintf("https://search.rakuten.co.jp/search/mall/%s/", keyword)
	searchCollector := rs.collector.Clone()

	// Update: Broader selector for search result items
	searchCollector.OnHTML("div.searchresultitem, div.dui-card.searchresultitem, .g-category-item", func(e *colly.HTMLElement) {
		if count >= maxProducts {
			return
		}

		// Update: More specific selectors for different page structures
		name := e.ChildText(".title, .g-category-item-name")

		// Try different price selectors
		priceText := e.ChildText(".important, .price, .g-category-item-price")
		price := extractPrice(priceText)

		// If price extraction failed, try data attribute
		if price == 0 {
			priceAttr := e.Attr("data-price")
			if priceAttr != "" {
				price = extractPrice(priceAttr)
			}
		}

		// Update: More selectors for product URL
		productURL := e.ChildAttr("a.title, a.g-category-item-name, a[data-url]", "href")
		if productURL == "" {
			productURL = e.ChildAttr("a", "href")
		}
		if productURL == "" {
			// Try data attribute
			productURL = e.Attr("data-url")
		}

		if productURL != "" && name != "" {
			id := extractProductID(productURL)
			product := models.NewProduct(id, name, productURL, "rakuten", price, "JPY")

			// Update: More selectors for image
			product.ImageURL = e.ChildAttr(".image img, .g-category-item-image img", "src")
			products = append(products, product)
			count++

			// Debug info
			log.Printf("Found product: %s, URL: %s, Price: %.2f", name, productURL, price)
		}
	})

	searchCollector.OnError(func(r *colly.Response, err error) {
		log.Printf("Error scraping search %s: %v", r.Request.URL, err)
	})

	// Start the search scraping
	searchCollector.Visit(searchURL)
	searchCollector.Wait()

	return products, nil
}

// extractPrice extracts a numerical price from text containing Japanese price formatting
func extractPrice(priceText string) float64 {
	// Clean up the price text
	priceText = strings.TrimSpace(priceText)
	priceText = strings.ReplaceAll(priceText, "円", "")
	priceText = strings.ReplaceAll(priceText, "¥", "")
	priceText = strings.ReplaceAll(priceText, ",", "")

	// Extract only the numbers using regex
	re := regexp.MustCompile(`\d+`)
	matches := re.FindString(priceText)

	if matches == "" {
		return 0
	}

	price, err := strconv.ParseFloat(matches, 64)
	if err != nil {
		log.Printf("Failed to parse price from text '%s' (extracted '%s'): %v", priceText, matches, err)
		return 0
	}

	return price
}

// Helper function to extract product ID from URL
func extractProductID(url string) string {
	// For Rakuten URLs like https://item.rakuten.co.jp/store/product-id/
	parts := strings.Split(url, "/")
	if len(parts) >= 5 {
		return parts[len(parts)-2]
	}

	// For books.rakuten.co.jp URLs like https://books.rakuten.co.jp/rb/14583459/
	if strings.Contains(url, "books.rakuten.co.jp") {
		parts := strings.Split(url, "/")
		for i, part := range parts {
			if part == "rb" && i+1 < len(parts) {
				return parts[i+1]
			}
		}
	}

	return fmt.Sprintf("rakuten-%d", time.Now().UnixNano())
}
