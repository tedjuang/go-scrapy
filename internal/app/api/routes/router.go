package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tedjuang/go-scrapy/internal/app/api/handlers"
	"github.com/tedjuang/go-scrapy/internal/app/api/middlewares"
)

// SetupRouter sets up the router with all API routes
func SetupRouter(dataDir string) (*gin.Engine, error) {
	r := gin.Default()

	// Add middleware
	r.Use(middlewares.Logger())

	// Create product handler
	handler, err := handlers.NewProductHandler(dataDir)
	if err != nil {
		return nil, err
	}

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.GET("", handler.GetAllProducts)
			products.GET("/:id", handler.GetProduct)
			products.POST("/scrape", handler.ScrapeProduct)
			products.POST("/search", handler.SearchProducts)
		}
	}

	// Serve OpenAPI documentation at a path that doesn't conflict with swagger UI
	r.StaticFile("/openapi/swagger.yaml", "./api/swagger.yaml")

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("/openapi/swagger.yaml")))

	// Static file handling for the UI (if we add a frontend)
	r.Static("/static", "./static")

	// Default route redirects to Swagger UI
	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/swagger/index.html")
	})

	return r, nil
}
