package main

import (
	"log"

	"github.com/damasdev/boring-go/app"
	"github.com/damasdev/boring-go/pkg/server"
)

func main() {
	// Create a new server instance
	s := server.New()

	// Load the application routes
	app.Boot(s)

	// Run the server
	if err := s.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
