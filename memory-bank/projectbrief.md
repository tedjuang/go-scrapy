# Scrapy - E-commerce Price Tracker

## Core Requirements

- Create a robust Go application for tracking prices on e-commerce websites
- Implement both command-line interface and REST API
- Allow scraping product details and tracking price changes over time
- Support searching for products by keyword
- Store product data locally in JSON format
- Start with Rakuten Japan support, design for extensibility

## Project Goals

- Provide users with a simple way to monitor price changes for products they're interested in
- Create an extensible architecture that can easily support additional e-commerce websites
- Deliver a clean REST API with Swagger documentation for integration with other systems
- Build a foundation that could be extended with additional features (price alerts, etc.)

## Project Scope

### In Scope

- Product detail scraping from supported e-commerce sites
- Price history tracking
- Product search functionality
- Local JSON data storage
- REST API with Swagger documentation
- Command-line interface
- Basic web UI for displaying tracked products

### Out of Scope (Future Enhancements)

- User authentication system
- Email/notification price alerts
- Advanced analytics
- Database storage (beyond JSON)
- Mobile applications

## Timeline

- Initial version with Rakuten Japan support
- REST API implementation
- Frontend web UI integration
- Support for additional e-commerce websites

## Success Criteria

- Successfully scrape and track products from Rakuten Japan
- Maintain price history for tracked products
- Provide functional search capability
- Deliver a working REST API with documentation
- Ensure good error handling and user experience
