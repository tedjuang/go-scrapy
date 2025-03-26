package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/tedjuang/go-scrapy/internal/models"
)

// Storage interface for product storage
type Storage interface {
	Save(product *models.Product) error
	GetAll() ([]*models.Product, error)
	GetByID(id string) (*models.Product, error)
}

// ProductStorage defines the interface for storing and retrieving products
type ProductStorage interface {
	Save(product *models.Product) error
	Get(id string) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	Delete(id string) error
}

// JSONFileStorage implements ProductStorage using a JSON file
type JSONFileStorage struct {
	filePath string
	products map[string]*models.Product
	mutex    sync.RWMutex
}

// NewJSONFileStorage creates a new JSON file storage
func NewJSONFileStorage(filePath string) (*JSONFileStorage, error) {
	storage := &JSONFileStorage{
		filePath: filePath,
		products: make(map[string]*models.Product),
	}

	// Load existing data if file exists
	if _, err := os.Stat(filePath); err == nil {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read storage file: %w", err)
		}

		var products []*models.Product
		if err := json.Unmarshal(data, &products); err != nil {
			return nil, fmt.Errorf("failed to parse storage file: %w", err)
		}

		for _, p := range products {
			storage.products[p.ID] = p
		}
	}

	return storage, nil
}

// Save stores a product in the storage
func (s *JSONFileStorage) Save(product *models.Product) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.products[product.ID] = product
	return s.writeToFile()
}

// Get retrieves a product by ID
func (s *JSONFileStorage) Get(id string) (*models.Product, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	product, exists := s.products[id]
	if !exists {
		return nil, fmt.Errorf("product with ID %s not found", id)
	}
	return product, nil
}

// GetAll returns all products
func (s *JSONFileStorage) GetAll() ([]*models.Product, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	products := make([]*models.Product, 0, len(s.products))
	for _, p := range s.products {
		products = append(products, p)
	}
	return products, nil
}

// Delete removes a product from storage
func (s *JSONFileStorage) Delete(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.products[id]; !exists {
		return fmt.Errorf("product with ID %s not found", id)
	}

	delete(s.products, id)
	return s.writeToFile()
}

// writeToFile persists the products to the JSON file
func (s *JSONFileStorage) writeToFile() error {
	products := make([]*models.Product, 0, len(s.products))
	for _, p := range s.products {
		products = append(products, p)
	}

	data, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal products: %w", err)
	}

	if err := os.WriteFile(s.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write to storage file: %w", err)
	}

	return nil
}

// GetByID returns a product by ID
func (s *JSONFileStorage) GetByID(id string) (*models.Product, error) {
	products, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}

	return nil, nil // Product not found, but no error
}
