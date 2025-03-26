# API Documentation

This directory contains documentation for the Scrapy API.

## Structure

- `swagger/` - Contains OpenAPI/Swagger specifications

  - `swagger.yaml` - The OpenAPI 3.0 specification for the API

- `decisions/` - Contains documentation about API design decisions

## Using the API Documentation

The Swagger specification can be viewed using any OpenAPI viewer, such as:

- Swagger UI (available at `/swagger/index.html` when the service is running)
- [Swagger Editor](https://editor.swagger.io/)
- [Redoc](https://redocly.github.io/redoc/)

## API Endpoints

The API provides the following main endpoints:

- `GET /api/v1/products` - Retrieve all products
- `GET /api/v1/products/{id}` - Get a specific product by ID
- `POST /api/v1/products/scrape` - Scrape a product from a URL
- `POST /api/v1/products/search` - Search for products on a website

For detailed documentation, please refer to the Swagger specification.
