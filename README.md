# Go-Scrapy - E-commerce Price Tracker

A Go application for tracking prices on e-commerce websites like Rakuten Japan.

## Features

- Scrape product details from supported e-commerce websites
- Track price changes over time
- Search for products by keyword
- Store product data locally in JSON format
- REST API with Swagger documentation

## Supported Websites

- Rakuten Japan (rakuten.co.jp)

## Installation

Make sure you have Go installed (version 1.18 or later), then:

```bash
# Clone the repository
git clone https://github.com/tedjuang/go-scrapy.git
cd go-scrapy

# Build the command line application
go build -o scrapy cmd/scrapy/main.go

# Build the API server
go build -o scrapy-api cmd/api/main.go
```

## Command Line Usage

```bash
# Track a specific product by URL
./scrapy -url "https://item.rakuten.co.jp/store/product-id/"

# Search for products
./scrapy -search "smartphone" -max 5

# Set a different data directory
./scrapy -url "https://item.rakuten.co.jp/store/product-id/" -data "./my-data"

# Show help
./scrapy -help
```

## API Server Usage

```bash
# Start the API server (with default development configuration)
./scrapy-api -env dev

# Start with production configuration
./scrapy-api -env prod

# Show help
./scrapy-api -help
```

Once started, you can access the Swagger UI at:
http://localhost:8080/swagger/index.html

### API Endpoints

- `GET /api/v1/products` - Get all tracked products
- `GET /api/v1/products/{id}` - Get a specific product by ID
- `POST /api/v1/products/scrape` - Scrape a product from a URL
- `POST /api/v1/products/search` - Search for products

## Command Line Arguments

### CLI Application

- `-url`: URL of the product to track
- `-search`: Search for products with this keyword
- `-website`: Website to scrape (default: "rakuten")
- `-max`: Maximum number of search results (default: 10)
- `-data`: Directory to store data (default: "./data")

### API Server

- `-env`: Environment to use (default: "dev")

## Data Storage

Product data is stored in a JSON file at `./data/products.json` (or the directory specified with the `-data` flag). Each time you track a new product or update an existing one, the price history is updated.

## Project Structure

The project follows the standard Go project layout:

```
go-scrapy/
├── api/                    # API specifications
├── cmd/                    # Application entry points
│   ├── api/                # API server
│   └── scrapy/             # CLI application
├── configs/                # Configuration files
│   ├── dev/                # Development environment configs
│   └── prod/               # Production environment configs
├── docs/                   # Documentation
│   ├── api/                # API documentation
│   │   ├── swagger/        # OpenAPI/Swagger specifications
│   │   └── decisions/      # API design decisions
│   ├── development/        # Developer documentation
│   └── usage/              # User documentation
├── internal/               # Private application code
│   ├── config/             # Configuration handling
│   ├── http/               # HTTP server implementation
│   │   ├── handlers/       # Request handlers
│   │   ├── middlewares/    # HTTP middlewares
│   │   └── routes/         # API routes
│   ├── models/             # Data models
│   ├── scraper/            # Scraper implementations
│   └── storage/            # Data storage
├── build/                  # Build and CI/CD configuration
├── deployments/            # Deployment configurations
├── scripts/                # Utility scripts
├── static/                 # Static web assets
├── test/                   # Additional test files
└── data/                   # Runtime data storage
```

## Extending

To add support for more websites:

1. Create a new scraper that implements the `scraper.Scraper` interface in `internal/scraper`
2. Register the scraper in the `scraper.NewScraperFactory()` function

## Documentation

- API documentation is available at `/swagger/index.html` when the server is running
- OpenAPI/Swagger specification is in `docs/api/swagger/swagger.yaml`
- Design decisions and rationale can be found in `docs/api/decisions/`

## License

MIT
