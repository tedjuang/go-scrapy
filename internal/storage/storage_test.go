package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tedjuang/go-scrapy/internal/models"
)

func TestJSONFileStorage(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "storage-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a test storage file
	testFilePath := filepath.Join(tmpDir, "test-products.json")

	// Initialize a new storage instance
	storage, err := NewJSONFileStorage(testFilePath)
	if err != nil {
		t.Fatalf("Failed to create storage: %v", err)
	}

	// Test saving a product
	testProduct := models.NewProduct(
		"test-123",
		"Test Product",
		"https://example.com/product",
		"rakuten",
		99.99,
		"JPY",
	)

	if err := storage.Save(testProduct); err != nil {
		t.Fatalf("Failed to save product: %v", err)
	}

	// Test retrieving the product
	retrievedProduct, err := storage.GetByID("test-123")
	if err != nil {
		t.Fatalf("Failed to get product: %v", err)
	}

	if retrievedProduct == nil {
		t.Fatal("Expected to get a product, got nil")
	}

	if retrievedProduct.ID != "test-123" {
		t.Errorf("Expected ID test-123, got %s", retrievedProduct.ID)
	}

	if retrievedProduct.Name != "Test Product" {
		t.Errorf("Expected Name 'Test Product', got %s", retrievedProduct.Name)
	}

	// Test saving a second product
	testProduct2 := models.NewProduct(
		"test-456",
		"Another Product",
		"https://example.com/another",
		"rakuten",
		199.99,
		"JPY",
	)

	if err := storage.Save(testProduct2); err != nil {
		t.Fatalf("Failed to save second product: %v", err)
	}

	// Test getting all products
	products, err := storage.GetAll()
	if err != nil {
		t.Fatalf("Failed to get all products: %v", err)
	}

	if len(products) != 2 {
		t.Errorf("Expected 2 products, got %d", len(products))
	}

	// Test loading storage from existing file
	newStorage, err := NewJSONFileStorage(testFilePath)
	if err != nil {
		t.Fatalf("Failed to create storage from existing file: %v", err)
	}

	loadedProducts, err := newStorage.GetAll()
	if err != nil {
		t.Fatalf("Failed to get all products from loaded storage: %v", err)
	}

	if len(loadedProducts) != 2 {
		t.Errorf("Expected 2 products from loaded storage, got %d", len(loadedProducts))
	}

	// Test product deletion
	if err := storage.Delete("test-123"); err != nil {
		t.Fatalf("Failed to delete product: %v", err)
	}

	// Verify product was deleted
	deletedProduct, err := storage.Get("test-123")
	if err == nil {
		t.Errorf("Expected error when getting deleted product, got nil")
	}

	if deletedProduct != nil {
		t.Errorf("Expected nil product after deletion, got %v", deletedProduct)
	}

	// Verify one product remains
	remainingProducts, err := storage.GetAll()
	if err != nil {
		t.Fatalf("Failed to get remaining products: %v", err)
	}

	if len(remainingProducts) != 1 {
		t.Errorf("Expected 1 remaining product, got %d", len(remainingProducts))
	}

	if remainingProducts[0].ID != "test-456" {
		t.Errorf("Expected remaining product ID test-456, got %s", remainingProducts[0].ID)
	}
}
