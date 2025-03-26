package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/tedjuang/go-scrapy/internal/models"
	"github.com/tedjuang/go-scrapy/internal/scraper"
	"github.com/tedjuang/go-scrapy/internal/storage"
)

// ProductHandler handles requests related to products
type ProductHandler struct {
	factory *scraper.ScraperFactory
	storage *storage.JSONFileStorage
}

// NewProductHandler creates a new product handler
func NewProductHandler(dataDir string) (*ProductHandler, error) {
	// Create storage
	store, err := storage.NewJSONFileStorage(filepath.Join(dataDir, "products.json"))
	if err != nil {
		return nil, err
	}

	// Create scraper factory
	factory := scraper.NewScraperFactory()

	return &ProductHandler{
		factory: factory,
		storage: store,
	}, nil
}

// ScrapeProductRequest represents a request to scrape a product
type ScrapeProductRequest struct {
	URL     string `json:"url" binding:"required" example:"https://item.rakuten.co.jp/book/14583459/"`
	Website string `json:"website" binding:"required" example:"rakuten"`
}

// SearchProductsRequest represents a request to search for products
type SearchProductsRequest struct {
	Keyword    string `json:"keyword" binding:"required" example:"nintendo switch"`
	Website    string `json:"website" binding:"required" example:"rakuten"`
	MaxResults int    `json:"max_results" example:"5"`
}

// ProductResponse represents the response for a product
type ProductResponse struct {
	Product *models.Product `json:"product"`
	Error   string          `json:"error,omitempty"`
}

// ProductsResponse represents the response for multiple products
type ProductsResponse struct {
	Products []*models.Product `json:"products"`
	Count    int               `json:"count"`
	Error    string            `json:"error,omitempty"`
}

// ScrapeProduct scrapes a product from a given URL
// @Summary Scrape a product from a URL
// @Description Scrape a product from a given URL and website
// @Tags products
// @Accept json
// @Produce json
// @Param request body ScrapeProductRequest true "Scrape Product Request"
// @Success 200 {object} ProductResponse "Product information"
// @Failure 400 {object} ProductResponse "Invalid request"
// @Failure 404 {object} ProductResponse "Scraper not found"
// @Failure 500 {object} ProductResponse "Server error"
// @Router /api/v1/products/scrape [post]
func (h *ProductHandler) ScrapeProduct(c *gin.Context) {
	var req ScrapeProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ProductResponse{
			Error: "Invalid request: " + err.Error(),
		})
		return
	}

	// Get the appropriate scraper
	s, exists := h.factory.GetScraper(req.Website)
	if !exists {
		c.JSON(http.StatusNotFound, ProductResponse{
			Error: "Scraper not found for website: " + req.Website,
		})
		return
	}

	// Scrape the product
	product, err := s.ScrapeProduct(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ProductResponse{
			Error: "Failed to scrape product: " + err.Error(),
		})
		return
	}

	// Save to storage
	if err := h.storage.Save(product); err != nil {
		c.JSON(http.StatusInternalServerError, ProductResponse{
			Error: "Failed to save product: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ProductResponse{
		Product: product,
	})
}

// SearchProducts searches for products
// @Summary Search for products
// @Description Search for products on a given website
// @Tags products
// @Accept json
// @Produce json
// @Param request body SearchProductsRequest true "Search Products Request"
// @Success 200 {object} ProductsResponse "Search results"
// @Failure 400 {object} ProductsResponse "Invalid request"
// @Failure 404 {object} ProductsResponse "Scraper not found"
// @Failure 500 {object} ProductsResponse "Server error"
// @Router /api/v1/products/search [post]
func (h *ProductHandler) SearchProducts(c *gin.Context) {
	var req SearchProductsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ProductsResponse{
			Error: "Invalid request: " + err.Error(),
		})
		return
	}

	// Default to 10 results if not specified
	if req.MaxResults <= 0 {
		req.MaxResults = 10
	}

	// Get the appropriate scraper
	s, exists := h.factory.GetScraper(req.Website)
	if !exists {
		c.JSON(http.StatusNotFound, ProductsResponse{
			Error: "Scraper not found for website: " + req.Website,
		})
		return
	}

	// Search for products
	products, err := s.ScrapeSearch(req.Keyword, req.MaxResults)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ProductsResponse{
			Error: "Failed to search for products: " + err.Error(),
		})
		return
	}

	// Save products to storage
	for _, p := range products {
		if err := h.storage.Save(p); err != nil {
			// Just log the error but continue
			// TODO: Add proper logging
		}
	}

	c.JSON(http.StatusOK, ProductsResponse{
		Products: products,
		Count:    len(products),
	})
}

// GetAllProducts returns all products
// @Summary Get all products
// @Description Get all stored products
// @Tags products
// @Produce json
// @Success 200 {object} ProductsResponse "All products"
// @Failure 500 {object} ProductsResponse "Server error"
// @Router /api/v1/products [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ProductsResponse{
			Error: "Failed to get products: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ProductsResponse{
		Products: products,
		Count:    len(products),
	})
}

// GetProduct returns a specific product by ID
// @Summary Get product by ID
// @Description Get a specific product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} ProductResponse "Product information"
// @Failure 404 {object} ProductResponse "Product not found"
// @Failure 500 {object} ProductResponse "Server error"
// @Router /api/v1/products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := h.storage.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ProductResponse{
			Error: "Failed to get product: " + err.Error(),
		})
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, ProductResponse{
			Error: "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, ProductResponse{
		Product: product,
	})
}
