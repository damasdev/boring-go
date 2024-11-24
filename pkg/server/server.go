package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/damasdev/boring-go/pkg/config"
	"github.com/gorilla/mux"
)

// Server is a struct that represents the server
type Server struct {
	router *mux.Router
	config *config.Server
}

// Option is a function that configures the server
type Option func(*config.Server)

// New initializes a new Server server instance with middleware
func New(opts ...Option) *Server {
	// Initialize server with default configuration
	s := &Server{
		config: defaultConfig(),
		router: mux.NewRouter(),
	}

	// Apply options
	for _, opt := range opts {
		opt(s.config)
	}

	// Return the server instance
	return s
}

// Router returns the server router
func (s *Server) Router() *mux.Router {
	return s.router
}

// Run starts the Server server
func (s *Server) Run() error {
	// Create a new server instance
	srv := &http.Server{
		Handler:      s.router,
		Addr:         fmt.Sprintf(":%s", s.config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server started on port %s", s.config.Port)

	// Start the server
	return srv.ListenAndServe()
}

// defaultConfig returns the default server configuration
func defaultConfig() *config.Server {
	return &config.Server{
		Port: "8000",
	}
}

// WithPort sets the port for the server configuration
func WithPort(port int) Option {
	return func(c *config.Server) {
		c.Port = fmt.Sprint(port)
	}
}
