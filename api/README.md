# API Specifications

This directory contains API entry points and infrastructure.

## API Documentation

The API documentation is located in the `/docs/api` directory:

- API specifications in OpenAPI/Swagger format
- API versioning documentation
- Design decisions

## Access API Documentation

When the service is running, Swagger UI is available at:

```
http://localhost:8080/swagger/index.html
```

## API Endpoints

The Scrapy API provides the following main endpoints:

- `GET /api/v1/products` - Retrieve all products
- `GET /api/v1/products/{id}` - Get a specific product by ID
- `POST /api/v1/products/scrape` - Scrape a product from a URL
- `POST /api/v1/products/search` - Search for products on a website

For detailed request/response specifications, please refer to the Swagger documentation.
