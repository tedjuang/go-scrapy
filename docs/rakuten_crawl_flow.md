# Rakuten Product Crawling Architecture and Data Flow

```mermaid
flowchart TD
    %% Main components
    Client[Client Application]
    APIServer[API Server]
    Handler[Product Handler]
    Factory[Scraper Factory]
    RakutenScraper[Rakuten Scraper]
    Colly[Colly Scraper Library]
    Storage[JSON File Storage]
    FSys[File System]

    %% External systems
    RakutenSite[Rakuten Website]

    %% Subgraphs
    subgraph "Client"
        Client
    end

    subgraph "API Layer"
        APIServer
        Handler
    end

    subgraph "Core Logic"
        Factory
        RakutenScraper
        Colly
    end

    subgraph "Persistence"
        Storage
        FSys
    end

    subgraph "External"
        RakutenSite
    end

    %% Data flow connections
    Client -->|POST /api/v1/products/scrape| APIServer
    APIServer -->|Route Request| Handler
    Handler -->|GetScraper("rakuten")| Factory
    Factory -->|Returns| RakutenScraper
    RakutenScraper -->|Uses| Colly
    Colly -->|HTTP Requests| RakutenSite
    RakutenSite -->|HTML Response| Colly
    Colly -->|Parsed Data| RakutenScraper
    RakutenScraper -->|Product Object| Handler
    Handler -->|Save Product| Storage
    Storage -->|Write JSON| FSys
    Storage -->|Saved Product| Handler
    Handler -->|Product Response| APIServer
    APIServer -->|JSON Response| Client

    %% Styles
    classDef apiStyle fill:#f9f,stroke:#333,stroke-width:2px;
    classDef coreStyle fill:#bbf,stroke:#333,stroke-width:2px;
    classDef persistStyle fill:#bfb,stroke:#333,stroke-width:2px;
    classDef externalStyle fill:#ddd,stroke:#333,stroke-width:2px;
    classDef clientStyle fill:#fbb,stroke:#333,stroke-width:2px;

    class APIServer,Handler apiStyle;
    class Factory,RakutenScraper,Colly coreStyle;
    class Storage,FSys persistStyle;
    class RakutenSite externalStyle;
    class Client clientStyle;
```

## Data Flow Description

1. **Request Initiation**:

   - Client sends a POST request to `/api/v1/products/scrape` with Rakuten product URL
   - API Server routes the request to the Product Handler

2. **Scraper Selection**:

   - Product Handler uses the Scraper Factory to get the appropriate scraper (Rakuten)
   - Scraper Factory returns the Rakuten Scraper implementation

3. **Web Scraping**:

   - Rakuten Scraper utilizes the Colly library for web scraping
   - Colly sends HTTP requests to the Rakuten website
   - Rakuten website returns HTML content
   - Colly parses the HTML and extracts product information
   - Rakuten Scraper transforms the parsed data into a Product object

4. **Data Persistence**:

   - Product Handler sends the Product object to JSON File Storage
   - Storage serializes the Product to JSON and writes to the file system
   - Storage confirms successful save to the Handler

5. **Response Delivery**:
   - Product Handler prepares a response with the Product information
   - API Server sends the JSON response back to the Client

This architecture follows clean separation of concerns with each component having a specific responsibility, making the system modular and maintainable.
