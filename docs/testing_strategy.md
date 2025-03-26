# Go-Scrapy Testing Strategy

## Overview

This document outlines the testing strategy for the Go-Scrapy project, including types of tests, priority order, and best practices.

## Test Types

### 1. Unit Tests

Unit tests focus on testing individual components in isolation.

**Priority Components for Unit Testing:**

1. ✅ **Models** (`internal/models`)

   - Test data structures and methods
   - Validate object creation and manipulation
   - Test edge cases like zero values and special characters

2. ✅ **Storage** (`internal/storage`)

   - Test CRUD operations
   - Test file persistence
   - Test error conditions
   - Test thread safety

3. ✅ **Scraper Factory** (`internal/scraper`)

   - Test factory creation
   - Test scraper registration and retrieval

4. ⏳ **API Handlers** (`internal/app/api/handlers`)
   - Test request validation
   - Test response generation
   - Test error handling

### 2. Integration Tests

Integration tests verify the interaction between multiple components.

**Priority for Integration Testing:**

1. **Scraper + Storage Integration**

   - Test the end-to-end flow of scraping and saving data

2. **API + Storage Integration**

   - Test API endpoints that read from and write to storage

3. **Complete Flow Integration**
   - Test the end-to-end flow from scraping to API response

### 3. Mocking Strategy

Effective testing requires appropriate mocking:

1. **External Dependencies**

   - Use a HTTP mocking library (like `httptest`) for scraper tests
   - Create mock implementations of interfaces for easier testing

2. **Interface-based Design**
   - Ensure all components use interfaces to enable mocking
   - Introduce test-friendly constructors that accept mock dependencies

## Test Implementation Roadmap

### Phase 1: Basic Unit Tests (Current)

- ✅ Unit tests for models
- ✅ Unit tests for storage
- ✅ Unit tests for scraper factory

### Phase 2: Enhanced Unit Tests

- 🔄 Refactor handlers to be more testable
- 🔄 Implement handler unit tests
- 🔄 Add test helpers for common tasks

### Phase 3: Integration Tests

- Add integration tests for key workflows
- Test scraping and storage combined operations
- Test API endpoints with real storage (but mocked scraping)

### Phase 4: End-to-End Tests

- Add end-to-end tests for key user journeys
- Test the CLI application
- Test the API server with real requests

## Best Practices

1. **Table-Driven Tests**

   - Use table-driven tests to test multiple scenarios
   - Group test cases logically by functionality

2. **Coverage Goals**

   - Aim for at least 80% code coverage
   - Focus on testing business logic and error handling

3. **Testing Error Conditions**

   - Test all error paths explicitly
   - Verify error messages and codes

4. **Test Independence**

   - Ensure tests can run independently
   - Clean up test data after each test

5. **Testing Performance**
   - Include benchmarks for performance-critical components
   - Test with realistic data volumes

## Tools

- Standard Go testing package
- httptest for testing HTTP handlers
- testify for more expressive assertions (optional)
- Coverage reports to measure test coverage

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test ./... -cover

# Run tests for a specific package
go test ./internal/models

# Run tests with verbose output
go test ./... -v
```

## Continuous Integration

In the future, we'll add GitHub Actions to:

- Run tests on each pull request
- Generate and publish coverage reports
- Run benchmarks to track performance

## Status & Next Steps

Currently, we have implemented unit tests for:

- ✅ Models
- ✅ Storage
- ✅ Scraper Factory

Next steps:

1. Refactor handlers to improve testability
2. Implement unit tests for handlers
3. Add integration tests for complete workflows
