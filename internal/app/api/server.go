// Package http provides the HTTP server implementation
package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/tedjuang/go-scrapy/internal/app/api/routes"
)

// Server represents the HTTP server
type Server struct {
	server  *http.Server
	dataDir string
}

// NewServer creates a new HTTP server
func NewServer(addr, dataDir string) *Server {
	return &Server{
		server: &http.Server{
			Addr: addr,
		},
		dataDir: dataDir,
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	router, err := routes.SetupRouter(s.dataDir)
	if err != nil {
		return err
	}

	s.server.Handler = router

	log.Printf("Starting server on %s\n", s.server.Addr)
	return s.server.ListenAndServe()
}

// Stop stops the HTTP server
func (s *Server) Stop(ctx context.Context) error {
	log.Println("Shutting down server...")
	return s.server.Shutdown(ctx)
}

// GracefulShutdown gracefully shuts down the server with a timeout
func (s *Server) GracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Stop(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v\n", err)
	}
	log.Println("Server exiting")
}
