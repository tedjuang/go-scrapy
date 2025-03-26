# Technical Context

## Technologies Used

### Core Technologies

- **Go (v1.24.1)**: Primary programming language
- **HTML/CSS/JavaScript**: Frontend web UI

### Dependencies

- **Colly**: Web scraping framework
  - Used for extracting data from e-commerce websites
  - Handles HTTP requests, HTML parsing, and data extraction
- **Gin**: Web framework for Go
  - Powers the REST API
  - Provides routing, middleware, and HTTP utilities
- **Swagger/OpenAPI**: API documentation standard
  - API documentation using OpenAPI 3.0 specification
  - Swagger UI for interactive API exploration
- **Goquery**: HTML parsing and manipulation

  - jQuery-like HTML manipulation for Go
  - Used within scrapers for DOM traversal

- **Go standard library**:
  - `net/http`: HTTP client and server
  - `encoding/json`: JSON serialization/deserialization
  - `flag`: Command-line flag parsing
  - `time`: Date and time utilities

### Development Tools

- **Go modules**: Dependency management
- **Makefile**: Build automation
- **Go testing**: Unit and integration testing

## Project Structure

Following the standard Go project layout as recommended by the community:

- `/api` - Contains ONLY API specifications (OpenAPI/Swagger, JSON schema files, protocol definition files)
- `/cmd` - Application entry points (API server and CLI)
- `/configs` - Configuration files for different environments
- `/docs` - Documentation and specifications
- `/internal` - Private implementation code
  - `/internal/app` - Application-specific code
    - `/internal/app/api` - API implementation and server
  - `/internal/models` - Data models
  - `/internal/scraper` - Scraper implementations
  - `/internal/storage` - Storage implementations
  - `/internal/config` - Configuration handling
- `/pkg` - Public libraries that can be imported by external applications
- `/build` - Build and packaging files
- `/deployments` - Deployment configuration
- `/scripts` - Utility scripts for development and operations
- `/test` - External test files and data

## Development Environment Setup

- Go 1.18+ installed
- Git for version control
- No external database required (uses local JSON storage)

## Build Process

1. Clone repository
2. Build CLI application: `go build -o scrapy cmd/scrapy/main.go`
3. Build API server: `go build -o scrapy-api cmd/api/main.go`

## Deployment

- Compiled binaries can be deployed on any platform supported by Go
- No special runtime requirements beyond the compiled binary
- Configuration via command-line flags and configuration files

## Technical Constraints

### Data Storage

- Currently limited to local JSON file storage
- Suitable for individual use cases
- May need migration to database for higher scale use

### Rate Limiting

- Must respect target websites' rate limits and robots.txt
- Implement appropriate delays between requests
- Consider using proxy rotation for high-volume scraping

### Website Changes

- Web scrapers are sensitive to website structure changes
- Need monitoring and maintenance to handle site updates
- Design scrapers for resilience and graceful failure

### Legal Considerations

- Comply with websites' terms of service
- Respect robots.txt directives
- Implement appropriate request headers (user agent)
- Consider ethical scraping practices

## Technical Extensibility Points

### Adding New Scrapers

1. Create new scraper that implements the `scraper.Scraper` interface
2. Register in `scraper.NewScraperFactory()`
3. No changes needed to existing code

### Storage Backend Alternatives

- Current implementation uses JSON file storage
- Design allows for alternative storage implementations
- Could be extended to support databases or cloud storage

### API Extensions

- API follows versioned design (`/api/v1/...`)
- New endpoints can be added without breaking compatibility
- Swagger documentation provides clear specification

## Development Guidelines

### Code Organization

- Follow standard Go project layout
- Keep implementation details in `/internal`
- Place application-specific code in `/internal/app`
- Place public interfaces in `/pkg` if needed
- Maintain clean separation of concerns
- `/api` directory must ONLY contain API specifications, not implementation code

### Documentation

- API specifications in OpenAPI/Swagger format in `/api`
- Design decisions documented in markdown files
- Comments for all exported functions and types
- Implementation details documented as needed

### Testing

- Unit tests for core functionality
- Integration tests for API endpoints
- Table-driven tests where appropriate
- Mock external dependencies when testing
