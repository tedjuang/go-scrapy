# System Patterns

## Architecture Overview

Scrapy follows the standard Go project layout with a modular architecture and clear separation of concerns:

```
scrapy/
├── api/                    # API specifications ONLY: OpenAPI/Swagger definitions, JSON schema files, protocol definition files
├── cmd/                    # Application entry points
│   ├── api/                # API server
│   └── scrapy/             # CLI application
├── configs/                # Configuration files
├── docs/                   # Documentation
│   ├── api/                # API documentation
│   │   └── decisions/      # API design decisions
│   ├── development/        # Developer documentation
│   └── usage/              # User documentation
├── internal/               # Private application code (not importable by external apps)
│   ├── app/                # Application-specific code
│   │   └── api/            # API implementation
│   │       ├── handlers/   # Request handlers
│   │       ├── middlewares/# HTTP middlewares
│   │       ├── routes/     # API routes
│   │       └── server.go   # API server implementation
│   ├── config/             # Configuration handling
│   ├── models/             # Data models
│   ├── scraper/            # Scraper implementations
│   └── storage/            # Data storage
├── pkg/                    # Public code that can be imported by other projects
├── build/                  # Build and CI/CD configuration
├── deployments/            # Deployment configurations
├── scripts/                # Utility scripts
├── static/                 # Static web assets
├── test/                   # Additional test files
└── data/                   # Runtime data storage
```

## Design Patterns

### Factory Pattern

- Used in `scraper.ScraperFactory` to create appropriate scraper instances
- Allows easy registration of new scrapers for different websites
- Centralizes scraper creation logic

### Interface-based Design

- `scraper.Scraper` interface defines the contract for all scrapers
- Enables adding new website scrapers without modifying existing code
- Promotes loose coupling between components

### Repository Pattern

- Storage package abstracts data persistence details
- Provides consistent interface for CRUD operations on product data
- Isolates data access logic from business logic

### Command Pattern

- CLI implementation uses command pattern for different operations
- Each flag/option represents a command with specific behavior
- Simplifies adding new commands/features

### Clean Architecture

- Core domain logic is independent of frameworks and delivery mechanisms
- Dependency injection is used to provide dependencies
- Follows the dependency rule: dependencies point inward

## Component Relationships

### API Layer Structure

1. Routes (`internal/app/api/routes`) define API endpoints
2. Middlewares (`internal/app/api/middlewares`) provide cross-cutting concerns
3. Handlers (`internal/app/api/handlers`) process requests and return responses
4. Core application logic is invoked from handlers

### Product Data Flow

1. User initiates request (CLI flag or API endpoint)
2. Appropriate scraper is selected via ScraperFactory
3. Scraper extracts product data from website
4. Data is converted to Product model
5. Storage component persists the data
6. Response is returned to user

### Search Data Flow

1. User provides search query
2. Query is passed to appropriate scraper
3. Scraper performs search on website
4. Results are converted to Product models
5. Results are returned to user (not stored by default)

## Technical Decisions

### Standard Go Project Layout

- Following the widely-accepted golang-standards/project-layout structure
- Clear separation between public (api, docs) and private code (internal)
- API specifications ONLY in `/api` directory - no implementation code
- Implementation code belongs in `/internal/app` or `/pkg` directories
- Application entry points in cmd directory
- Documentation and specifications in their designated directories

### API Documentation Approach

- OpenAPI/Swagger specifications in `/api` directory
- API design decisions documented in `docs/api/decisions`
- Runtime documentation served through the API server

### Go Language

- Excellent for concurrent operations (web scraping)
- Strong standard library with HTTP support
- Good performance characteristics
- Cross-platform compatibility

### Local JSON Storage

- Simple and portable storage solution
- No external database dependencies
- Suitable for the expected data volume
- Easy backup and version control

### Colly Web Scraping Framework

- Powerful Go scraping library with good performance
- Handles common scraping tasks (session management, rate limiting)
- Extensible for custom scraping needs

### Gin Web Framework (API)

- Lightweight and fast HTTP router
- Good middleware support
- Excellent for REST API implementation

### API Versioning Strategy

- URI path versioning (`/api/v1/...`)
- Allows backward compatibility while evolving the API
- Clear documentation of versioning decisions
