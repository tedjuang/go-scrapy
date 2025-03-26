# Active Context

## Current Work Focus

- Project structure refinement to conform to standard Go project layout
- Restructuring internal code to follow app/pkg pattern
- Ensuring `/api` directory contains ONLY API specifications (OpenAPI/Swagger, JSON schema) and no implementation code
- Removing duplicated files and code
- Fixing API server routing issues

## Recent Changes

- Restructured project to follow golang-standards/project-layout guidelines:
  - Moved API implementation code from `/internal/http` to `/internal/app/api`
  - Removed duplicate implementations and files
  - Organized code into more semantic directories (app vs http)
  - Placed API specifications (OpenAPI/Swagger) in `/api` directory
  - Created proper documentation structure in `/docs`
  - Added API design decisions documentation
- Fixed Swagger route conflicts in API server
  - Resolved issues with wildcard routes and specific file routes
  - Updated OpenAPI specification path
- Created clear separation between public interfaces and private implementation
- Enhanced project documentation structure
- Ensured `/api` directory only contains specification files, not Go implementation code

## Next Steps

### Short-term Tasks

1. Complete implementation of remaining API endpoints
2. Add comprehensive test coverage for API handlers
3. Implement structured logging throughout the application
4. Enhance error handling and validation
5. Implement proper configuration management

### Medium-term Tasks

1. Add support for additional e-commerce websites
2. Implement price alert functionality
3. Enhance search capabilities with filtering options
4. Add data export functionality (CSV, Excel)
5. Improve scraper resilience to website changes

### Long-term Vision

1. Consider database storage options for scalability
2. Implement user accounts and access control
3. Add notification system for price alerts
4. Develop mobile applications
5. Consider cloud deployment and hosting options

## Active Decisions and Considerations

### Technical Decisions

- Adopted standard Go project layout for better maintainability and alignment with community standards
- Moved implementation code to `/internal/app` following semantic naming over technical naming
- Avoided using "http" as a directory name as it's a protocol name rather than a functional area
- Separated API specifications into `/api` directory, ensuring it only contains specification files and not Go code
- Using the Gin web framework for API implementation
- Keeping business logic separate from transport layer

### Open Questions

- How to effectively maintain OpenAPI specifications in the `/api` directory as the API evolves
- Best approach for handling website structure changes in scrapers
- Strategies for implementing comprehensive tests for scrapers
- Options for implementing price alerts and notifications

### Current Challenges

- Ensuring all code follows the new project structure
- Making sure import paths are updated correctly after restructuring
- Resolving routing conflicts in Gin (specifically with Swagger UI and static files)
- Designing a cohesive and consistent API
- Balancing scraping frequency with website rate limits
- Planning for extensibility as new websites are added

## Current Development Status

- Project structure has been refactored to follow standard Go layout with app/pkg pattern
- API specifications and implementation are properly separated
  - Specifications in `/api`
  - Implementation in `/internal/app/api`
- Swagger UI is working correctly with the API server
- Documentation structure is in place
- Core scraping functionality works for Rakuten Japan
- API is functional but needs refinement and testing
- Storage system is implemented but basic
