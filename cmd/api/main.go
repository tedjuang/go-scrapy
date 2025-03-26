package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	http "github.com/tedjuang/go-scrapy/internal/app/api"
	"github.com/tedjuang/go-scrapy/internal/config"
)

func main() {
	// Parse flags
	env := flag.String("env", "dev", "Environment (dev, prod)")
	flag.Parse()

	// Load configuration
	cfg, err := config.LoadConfig(*env)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create server address from config
	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	// Create server
	server := http.NewServer(serverAddr, cfg.Data.Dir)

	// Start server in a goroutine
	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Server started on %s", serverAddr)
	log.Printf("Swagger UI available at http://%s/swagger/index.html", serverAddr)

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Gracefully shutdown
	server.GracefulShutdown()
}
