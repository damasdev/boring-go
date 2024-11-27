package app

import (
	"net/http"
	"strings"

	"github.com/damasdev/boring-go/app/landing"
	"github.com/damasdev/boring-go/app/user"
	"github.com/damasdev/boring-go/pkg/middleware"
	"github.com/damasdev/boring-go/pkg/server"
)

// Load the application.
func Load(s *server.Server) {
	// Get the router from the server.
	r := s.Router()

	// Middleware
	r.Use(middleware.Recover())

	// Serve static files from the public directory.
	fileServer := http.FileServer(http.Dir("public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public", restrictedFileServer(fileServer)))

	// Define the routes.
	landing.Routes(r)
	user.Routes(r)
}

// restrictedFileServer is a middleware that prevents directory listing.
func restrictedFileServer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
